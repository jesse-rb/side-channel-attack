package main

import (
	"fmt"
	"time"
	"log"
	"math"
)

func experiment(bytes []byte, steps int, incr int) (time.Duration) {
	var lengthMod int = len(bytes)-1
	var elapsedAvg int64 = 0
	var reps int = 50;
	for i := 0; i < reps+1; i++ {
		begin := time.Now()
		for j := 0; j < steps; j++ {
			bytes[(j*incr)&lengthMod]++ // Fast modulus operation
		}
		if i == 0 { continue }
		elapsedAvg += time.Since(begin).Nanoseconds()
	}	
	
	elapsedAvg = elapsedAvg/int64(reps)
	return time.Duration(elapsedAvg)
}

func main() {
	fmt.Printf("Time (Nanoseconds),Array Size (KiB), Log 2 Array Size (B),St eps,Increment\n")
	var arrSize = 1024; // Start at 1 KB
	for i := 1; i <= 21; i++ {
		var bytes	[]byte	= make([]byte, arrSize)
		var steps	int 	= 64 * 1024 * 1024
		var incr	int		= 16
		elapsed := experiment(bytes, steps, incr)
		arrSize = arrSize*2 // Double array size each time
		log.Println(fmt.Sprintf("%d", i))
		fmt.Printf("%d,%d,%f,%d,%d\n", elapsed.Nanoseconds(), arrSize/1024, math.Log2(float64(arrSize)), steps, incr)
	}
}