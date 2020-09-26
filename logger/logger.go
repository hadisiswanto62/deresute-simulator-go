package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

const LOG_DIR = "F:/Temporary/2020.09.26 - deresute-simulator log"

func Log(text string) {
	x := instance()
	x.Println(text)
}

func Logf(text string, a ...interface{}) {
	Log(fmt.Sprintf(text, a...))
}

func instance() *log.Logger {
	if logger == nil {
		time := time.Now().Unix()
		logFile := fmt.Sprintf("%s/%d.txt", LOG_DIR, time)
		f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		logger = log.New(f, "prefix here", log.LstdFlags)
	}
	return logger
}

var logger *log.Logger
