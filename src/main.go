package main

import (
	"fmt"
	dnnlparser "parser"
	"strconv"
)

func main() {
	layerinfo := dnnlparser.ParseVerbose("../verbose_example_output")
	for i := range layerinfo {
		fmt.Println("layer " + strconv.Itoa(i+1))
		fmt.Println("Operation name: " + layerinfo[i].OpName)
		fmt.Println("Execution time: " + layerinfo[i].ExeTime)
		fmt.Println("Data format: " + layerinfo[i].DataFormat)
		fmt.Println("Propogation type" + layerinfo[i].Propogation)
		fmt.Println("Metadata: " + layerinfo[i].MetaData)
	}
}
