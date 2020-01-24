package dnnl

import "time"

var (
	startTime time.Time
)

func init() {
	startTime = time.Now()
}
