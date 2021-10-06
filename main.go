package main

import (
	"fmt"
	"time"
)

func experiment(bytes []byte, steps int, incr int) (time.Duration) {
	begin := time.Now()

	var index int = 0
	for i := 0; i < len(bytes)-1; i++ {
		if index < len(bytes)-1 { _ = bytes[index] }
		index += incr
		if index >= len(bytes) { index = 0 }
	}

	elapsed := time.Since(begin)
	return elapsed
}

func average(bytes []byte, steps int, incr int) (time.Duration) {
	var nanoSecondsAvg int64 = 0;
	var iterations int64 = 50
	var i int64 = 0
	for i = 0; i < iterations; i++ {
		nanoSecondsAvg += experiment(bytes, steps, incr).Nanoseconds()
	}
	return (time.Duration(nanoSecondsAvg/iterations))
}

func main() {
	fmt.Printf("Time (Microseconds),Array Size (Bytes),Steps,Increment\n")
	for i := 0; i < 100; i++ {
		var bytes		[]byte	= make([]byte, i*1024)
		var steps	int 	= i*1024
		var incr	int		= i*16	
		elapsed := average(bytes, steps, incr)
		fmt.Printf("%d,%d,%d,%d\n", elapsed.Microseconds(), len(bytes), steps, incr)
	}
}