package dnnl

import (
	"fmt"
	"context"
	"os"
	"strconv"
	"time"

	"github.com/Unknwon/com"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/rai-project/tracer"
)

// Verbose mode const
// A valid sample output line is given as:
// mkldnn_verbose,exec,reorder,jit:uni,undef,in:f32_oihw out:f32_Ohwi16o,num:1,64x3x11x11,0.0219727
var (
	DNNLMARKER string = "mkldnn_verbose"
	OP         string = "exec"
	INFO       string = "info"
)

// layer 46
// Operation name: inner_product
// Execution time: 3.75903
// Data format: fsrc:nc fwei:oi fbia:x fdst:nc
// Propogation typeforward_inference
// Metadata: gemm:blas, mb1ic9216oc4096
type TraceEvent struct {
	OpName      string // primitive kind
	Timestamp   int64
	ExeTime     time.Duration
	DataFormat  string            // input/output data format
	Propogation string            // propogation
	MetaData    map[string]string // meta data
}

type TraceEvents []*TraceEvent

func (t TraceEvents) Len() int           { return len(t) }
func (t TraceEvents) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t TraceEvents) Less(i, j int) bool { return t[i].Timestamp < t[j].Timestamp }

type Trace struct {
	TraceEvents TraceEvents
	filename    string
	StartTime   time.Time
	EndTime     time.Time
}

func (t Trace) Len() int           { return t.TraceEvents.Len() }
func (t Trace) Swap(i, j int)      { t.TraceEvents.Swap(i, j) }
func (t Trace) Less(i, j int) bool { return t.TraceEvents.Less(i, j) }

func NewProfile(tmpDir string) (*Trace, error) {
	if !com.IsDir(tmpDir) {
		os.MkdirAll(tmpDir, os.FileMode(0755))
	}
	filename, err := tempFile(tmpDir, "profile-", ".out")
	if err != nil {
		return nil, errors.Errorf("cannot create temporary file in %v", tmpDir)
	}
	verboseOut := ParseVerbose("/Users/apple/go-mkl-dnn/example/verbose_example_output")
	fmt.Println(filename)
	fmt.Println(startTime)
	return &Trace{
		TraceEvents: verboseOut,
		filename:    filename,
		StartTime:   startTime,
		EndTime:     startTime,
	}, nil
}

func (t *Trace) Publish(ctx context.Context, lvl tracer.Level, opts ...opentracing.StartSpanOption) error {

	accumTime := t.StartTime
	layerIdx := 0

	for _, event := range t.TraceEvents {
		tags := opentracing.Tags{}
		for k, v := range event.MetaData {
			tags[k] = v
			//println(k, v)
		}
		tags["layer idx"] = strconv.Itoa(layerIdx)
		s, _ := tracer.StartSpanFromContext(
			ctx,
			lvl,
			event.OpName,
			opentracing.StartTime(accumTime),
			tags,
		)

		if s == nil {
			continue
		}

		accumTime.Add(event.ExeTime)

		// duration := endTime.Sub(startEntry.startTime).Nanoseconds()
		s.
			// SetTag("end_timestamp", timeUnit*time.Duration(event.Timestamp)).
			// SetTag("endtime", endTime).
			// SetTag("duration(ns)", duration).
			FinishWithOptions(opentracing.FinishOptions{
				FinishTime: accumTime,
			})
	}

	t.EndTime = accumTime

	return nil
}
