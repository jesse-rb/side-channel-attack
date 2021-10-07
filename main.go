package main

import (
	"fmt"
	"time"
	"log"
)

func experiment(bytes []byte, steps int, incr int) (time.Duration) {
	var nanoSecondsAvg int64 = 0;
	var lengthMod int = len(bytes)-1
	for i := 0; i < len(bytes)-1; i++ {
		begin := time.Now()
		bytes[(i*incr)&lengthMod]++
		elapsed := time.Since(begin)
		nanoSecondsAvg += elapsed.Nanoseconds()
	}

	var total int = len(bytes)
	nanoSecondsAvg = nanoSecondsAvg/(int64(total))
	return time.Duration(nanoSecondsAvg)
}

func main() {
	fmt.Printf("Time (Nanoseconds),Array Size (Bytes),St eps,Increment\n")
	for i := 1; i <= 100; i++ {
		var bytes	[]byte	= make([]byte, i*1024*1024)
		var steps	int 	= i*1024*1024
		var incr	int		= 16
		elapsed := experiment(bytes, steps, incr)
		log.Println(fmt.Sprintf("%d", i));
		fmt.Printf("%d,%d,%d,%d\n", elapsed.Nanoseconds(), len(bytes), steps, incr)
	}
}