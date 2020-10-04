package stopwatch

import (
	"fmt"
	"time"
)

var start int64
var end int64

func Start() {
	start = time.Now().UnixNano()
}

func Stop() {
	end = time.Now().UnixNano()
}

func Print() {
	fmt.Printf("Took %v seconds\n", float64(end-start)/float64(time.Second))
}
