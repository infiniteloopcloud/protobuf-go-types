package log

import (
	"fmt"
	"os"
)

const (
	logfilePathEnv = "LOGFILE_PATH"
	logEnabled     = "DEBUG"
)

var logFilePath = os.Getenv(logfilePathEnv)

var logFile *os.File

func init() {
	if os.Getenv(logEnabled) != "true" {
		return
	}
	var err error
	if logFilePath == "" {
		logFilePath = "/tmp/proto.log"
	}
	logFile, err = os.Create(logFilePath)
	if err != nil {
		panic(err)
	}
}

func Log(format string, args ...any) {
	if os.Getenv(logEnabled) != "true" {
		return
	}
	fmt.Fprintf(logFile, format+"\n", args...)
}
