package dnnl

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"regexp"
	"time"
)

/* ParseVerbose: parse the dnnl verbose mode output text file
 * input: filename
 * output: slice of TraceEvent struct
 * metadata description: ic, oc	Input/Output channels (aka feature maps)
 * id, ih, iw	Input depth, height and width
 * od, oh, ow	Output depth, height and width
 * kd, kh, kw	Kernel (filter, weights) depth, height and width
 * sd, sh, sw	Convolution stride over depth, height and width
 * pd, ph, pw	Convolution front, top and left padding
 * mb	Minibatch (amount of images processed at once)
 * g Groups (a way to reduce the amount of computations, see Alexnet topology)
 * example: mb1_ic3oc64_ih224oh55kh11sh4dh0ph2_iw224ow55kw11sw4dw0pw2
 */

func ParseVerbose(fName string) TraceEvents {
	f, err := os.Open(fName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	verboseOut := TraceEvents{}

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.Split(s.Text(), ",")
		if strings.Compare(line[0], DNNLMARKER) == 0 && strings.Compare(line[1], INFO) == 0 {
			fmt.Println("verbose mode header, start to parse verbose output...")
			continue
		} else if strings.Compare(line[0], DNNLMARKER) == 0 && strings.Compare(line[1], OP) == 0 {
			layer := new(TraceEvent)
			layer.OpName = line[2]
			//f, _ := strconv.ParseFloat(line[len(line)-1], 64)
			layer.ExeTime, _ = time.ParseDuration(line[len(line)-1]+"ms")
			layer.Propogation = line[4]
			layer.DataFormat = line[5]
			meta := make(map[string]string)
			// if len(line[3]) != 0 {
			// 	meta["jit"] = strings.Split(line[3], ":")[1]
			// }
			if len(line[6]) != 0 {
				temp := strings.Split(line[6], ":")
				meta[temp[0]] = temp[1]
			}
			if len(line[7]) != 0 {
				match, _ := regexp.MatchString("([0-9]+)x", line[7])
				if(match) {
					meta["dimension"] = line[7]
				} else {
					dim := strings.Join(strings.Split(line[7], "_"), "")
					re := regexp.MustCompile("\\D+|\\d+")
					details := re.FindAllString(dim, -1)
					i := 0
					for i < len(details) {
						meta[details[i]] = details[i+1]
						i += 2
					}
				}
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


