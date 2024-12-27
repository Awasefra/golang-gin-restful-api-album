package utils

import (
	"log"
	"os"
)

var logger *log.Logger

func init() {
	// Open or create the log file
	logFile, err := os.OpenFile("logger.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}

	// Set the logger to write to both file and console
	logger = log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)
}

// Info logs informational messages
func Info(message string) {
	logger.SetPrefix("[INFO] ")
	logger.Println(message)
}

// Error logs error messages
func Error(message string) {
	logger.SetPrefix("[ERROR] ")
	logger.Println(message)
}

// Warning logs warning messages
func Warning(message string) {
	logger.SetPrefix("[WARNING] ")
	logger.Println(message)
}
