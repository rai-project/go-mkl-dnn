package main

import (
	"fmt"
	//"strconv"
	"context"
	dnnl "github.com/rai-project/go-mkl-dnn"
	"github.com/rai-project/tracer"
)

func main() {
	defer tracer.Close()
	//layerinfo := dnnl.ParseVerbose("verbose_example_output")
	//fmt.Println("in main function")
	trace, _ := dnnl.NewProfile("/Users/apple/go-mkl-dnn/example/")
	// for t := range trace.TraceEvents {
	// 	println(trace.TraceEvents[t].OpName)
	// 	println(trace.TraceEvents[t].ExeTime)
	// }
	//fmt.Println(trace.TraceEvents.Len())
	ctx := context.Background()
	e := trace.Publish(ctx,tracer.FULL_TRACE, )
	if e != nil {
		fmt.Println("Failed to publish tracer")
	}
	
	// for i := range layerinfo {
	// 	fmt.Println("layer " + strconv.Itoa(i+1))
	// 	// fmt.Println("Operation name: " + layerinfo[i].OpName)
	// 	// fmt.Println("Execution time: " + layerinfo[i].ExeTime)
	// 	// fmt.Println("Data format: " + layerinfo[i].DataFormat)
	// 	// fmt.Println("Propogation type" + layerinfo[i].Propogation)
	// 	// fmt.Println("Metadata: " + layerinfo[i].MetaData)
	// 	for key, value := range layerinfo[i].MetaData {
	// 		fmt.Println(key, value)
	// 	}
	// }
}
// import (
// 	"fmt"
// 	"os"
// 	dnnl "github.com/rai-project/go-mkl-dnn"
// 	"github.com/opentracing/opentracing-go/log"
// )
// func main() {
// 	if len(os.Args) != 2 {
// 		panic("ERROR: Expecting one argument")
// 	}

// 	tracer, closer := dnnl.InitJaeger("hello-world")
// 	defer closer.Close()

// 	helloTo := os.Args[1]

// 	span := tracer.StartSpan("say-hello")
// 	span.SetTag("hello-to", helloTo)

// 	helloStr := fmt.Sprintf("Hello, %s!", helloTo)
// 	span.LogFields(
// 		log.String("event", "string-format"),
// 		log.String("value", helloStr),
// 	)

// 	println(helloStr)
// 	span.LogKV("event", "println")

// 	span.Finish()
// }