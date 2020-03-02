package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
	//"strconv"
	"context"

	"github.com/rai-project/config"
	dnnl "github.com/rai-project/go-mkl-dnn"
	"github.com/rai-project/tracer"
	_ "github.com/rai-project/tracer/all"
	opentracing "github.com/opentracing/opentracing-go"
)

var configOptions = []config.Option{
	config.AppName("carml"),
	config.ColorMode(true),
	config.DebugMode(true),
	config.VerboseMode(true),
}

func main() {
	config.Init(configOptions...)
	defer tracer.Close()
	//layerinfo := dnnl.ParseVerbose("verbose_example_output")
	//fmt.Println("in main function")
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_ = pwd
	trace, _ := dnnl.NewProfile(filepath.Join(pwd, "verbose_example_output"))
	// for t := range trace.TraceEvents {
	// 	println(trace.TraceEvents[t].OpName)
	// 	println(trace.TraceEvents[t].ExeTime)
	// }
	//fmt.Println(trace.TraceEvents.Len())
	//ctx := context.Background()
	t := time.Now()
	span, ctx := tracer.StartSpanFromContext(context.Background(), tracer.MODEL_TRACE, "alexnet", opentracing.StartTime(t))
	acc, e := trace.Publish(t, ctx, tracer.FULL_TRACE)
	if e != nil {
		fmt.Println("Failed to publish tracer")
	}
	span.
	SetTag("endtime", acc).
	FinishWithOptions(opentracing.FinishOptions{
		FinishTime: acc,
	})	
	//defer span.Finish()
	
	// pp.Println(span)
	//_ = span
	//_ = trace
	//dnnl.UpdateStartTime()
	
	
	// span.
	// SetTag("endtime", acc).
	// FinishWithOptions(opentracing.FinishOptions{
	// 	FinishTime: acc,
	// })
	
	//span.Finish()
}

