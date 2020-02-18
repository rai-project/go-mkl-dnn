package dnnl

import "time"

var (
	startTime time.Time
)

func UpdateStartTime() {
	startTime = time.Now()

}
