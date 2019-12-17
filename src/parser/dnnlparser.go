package dnnlparser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*  This package is used for parsing DNNL verbose mode output
 *  A valid sample output line is given as: mkldnn_verbose,exec,reorder,jit:uni,undef,in:f32_oihw out:f32_Ohwi16o,num:1,64x3x11x11,0.0219727
 */

// Verbose mode const
const DNNLMARKER string = "mkldnn_verbose"
const OP string = "exec"
const INFO string = "info"

type DnnlVerbose struct {
	OpName      string // primitive kind
	ExeTime     string // execution time in millisecond
	DataFormat  string // input/output data format
	Propogation string // propogation
	MetaData    string // meta data
}

/* ParseVerbose: parse the dnnl verbose mode output text file
 * input: filename
 * output: slice of dnnlVerbose struct
 */

func ParseVerbose(fName string) []*DnnlVerbose {
	f, err := os.Open(fName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	verboseOut := []*DnnlVerbose{}

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.Split(s.Text(), ",")
		if strings.Compare(line[0], DNNLMARKER) == 0 && strings.Compare(line[1], INFO) == 0 {
			fmt.Println("verbose mode header, start to parse verbose output...")
			continue
		} else if strings.Compare(line[0], DNNLMARKER) == 0 && strings.Compare(line[1], OP) == 0 {
			layer := new(DnnlVerbose)
			layer.OpName = line[2]
			layer.ExeTime = line[len(line)-1]
			layer.Propogation = line[4]
			layer.DataFormat = line[5]
			meta := ""
			if len(line[3]) != 0 {
				meta += line[3] + ", "
			}
			if len(line[6]) != 0 {
				meta += line[6] + ", "
			}
			if len(line[7]) != 0 {
				meta += line[7]
			}
			layer.MetaData = meta
			verboseOut = append(verboseOut, layer)
		} else {
			break
		}
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
	return verboseOut

}
