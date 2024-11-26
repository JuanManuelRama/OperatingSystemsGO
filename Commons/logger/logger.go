package logger

import (
	"log"
)

const (
	reset  = "\033[0m"  // Reset to default color
	yellow = "\033[33m" // Yellow color for warnings
	green  = "\033[32m" // Green color for success
	red    = "\033[31m" // Red color for errors
)

func Warning(message string) {
	log.Printf("%sWARNING: %s%s", yellow, message, reset)
}

func Error(message string) {
	log.Fatalf("%sERROR: %s%s", red, message, reset)
}

func Info(message string) {
	log.Println("INFO:", message)
}
