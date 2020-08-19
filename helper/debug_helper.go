package helper

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"time"

	"github.com/hadisiswanto62/deresute-simulator-go/config"
)

// MeasureTime is used to log execution time of a function.
// How to use: put `defer helper.MeasureTime(time.Now(), func_name_here`
// on the first line of the function
func MeasureTime(init time.Time, name string) {
	elapsed := time.Since(init)
	log.Printf("%s took %f s", name, elapsed.Seconds())
}

func GetSkillAlwaysActive() bool {
	return config.SkillAlwaysActive
}

var startTime int64

// MeasureTimeSincePrev returns the duration from the previous MeasureTimeSincePrev call (first call will return -1)
func MeasureTimeSincePrev() float64 {
	if startTime == 0 {
		// log.Printf("Initialized")
		startTime = time.Now().UnixNano()
		return -1
	} else {
		now := time.Now().UnixNano()
		elapsed := now - startTime
		// log.Printf("%s took %fs", name, float64(elapsed)/1000000000.0)
		startTime = now
		return float64(elapsed) / 1000000000.0
	}
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
