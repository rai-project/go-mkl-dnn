package dnnl

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strings"
// )

// /* ParseVerbose: parse the dnnl verbose mode output text file
//  * input: filename
//  * output: slice of TraceEvent struct
//  */
// func ParseVerbose(fName string) TraceEvents {
// 	f, err := os.Open(fName)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()

// 	verboseOut := TraceEvents{}

// 	s := bufio.NewScanner(f)
// 	for s.Scan() {
// 		line := strings.Split(s.Text(), ",")
// 		if strings.Compare(line[0], DNNLMARKER) == 0 && strings.Compare(line[1], INFO) == 0 {
// 			fmt.Println("verbose mode header, start to parse verbose output...")
// 			continue
// 		} else if strings.Compare(line[0], DNNLMARKER) == 0 && strings.Compare(line[1], OP) == 0 {
// 			layer := new(TraceEvent)
// 			layer.OpName = line[2]
// 			layer.ExeTime = line[len(line)-1]
// 			layer.Propogation = line[4]
// 			layer.DataFormat = line[5]
// 			meta := ""
// 			if len(line[3]) != 0 {
// 				meta += line[3] + ", "
// 			}
// 			if len(line[6]) != 0 {
// 				meta += line[6] + ", "
// 			}
// 			if len(line[7]) != 0 {
// 				meta += line[7]
// 			}
// 			layer.MetaData = meta
// 			verboseOut = append(verboseOut, layer)
// 		} else {
// 			break
// 		}
// 	}
// 	err = s.Err()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return verboseOut
// }
