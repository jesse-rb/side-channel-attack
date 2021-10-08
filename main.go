package main

import (
	"fmt"
	"time"
	"log"
)

func experiment(bytes []byte, steps int, incr int) (time.Duration) {
	var lengthMod int = len(bytes)-1
	var elapsedAvg int64 = 0

	for i := 0; i <= 50; i++ {
		begin := time.Now()
		for i := 0; i < len(bytes)-1; i++ {
			bytes[(i*incr)&lengthMod]++ // Fast modulus operation
			// log.Println(fmt.Sprintf("%d", bytes[(i*incr)&lengthMod]))
		}
		elapsedAvg += time.Since(begin).Nanoseconds()
	}
	
	elapsedAvg = elapsedAvg/int64(len(bytes))
	return time.Duration(elapsedAvg)
}

func main() {
	fmt.Printf("Time (Nanoseconds),Array Size (MegaBytes),St eps,Increment\n")
	for i := 1; i <= 32; i++ {
		var bytes	[]byte	= make([]byte, i*1024*1024)
		var steps	int 	= 64 * 1024 * 1024
		var incr	int		= 16
		elapsed := experiment(bytes, steps, incr)
		log.Println(fmt.Sprintf("%d", i));
		fmt.Printf("%d,%d,%d,%d\n", elapsed.Nanoseconds(), len(bytes)/(1024*1024), steps, incr)
	}
}