package helper

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"time"
)

// MeasureTime is used to log execution time of a function.
// How to use: put `defer helper.MeasureTime(time.Now(), func_name_here`
// on the first line of the function
func MeasureTime(init time.Time, name string) {
	elapsed := time.Since(init)
	log.Printf("%s took %f s", name, elapsed.Seconds())
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Println("-----")
	fmt.Printf("Alloc = %v KB\n", bToKb(m.Alloc))
	fmt.Printf("TotalAlloc = %v KB\n", bToKb(m.TotalAlloc))
	fmt.Printf("Sys = %v KB\n", bToKb(m.Sys))
	fmt.Printf("NumGC = %v\n", m.NumGC)
	fmt.Println("-----")
}

func bToKb(b uint64) uint64 {
	return b / 1024
}

// RandInt returns a random integer between min and max-1 [min, max)
func RandInt(min, max int) int {
	roll := rand.Intn(max - min)
	return roll + min
}
