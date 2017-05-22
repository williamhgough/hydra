package logger

import (
	"log"
	"os"
	"sync"
)

// Logger - Custom Logger Structure
type Logger struct {
	*log.Logger
	filename string
}

// Loggy - Instance of Logger
var loggy *Logger
var once sync.Once

// GetInstance - Create a singleton instance of the custom logger.
func GetInstance() *Logger {
	once.Do(func() {
		loggy = createLogger("loggy.log")
	})
	return loggy
}

func createLogger(name string) *Logger {
	file, _ := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	return &Logger{
		filename: name,
		Logger:   log.New(file, "Hydra-", log.Lshortfile),
	}
}

// Rec - record new message to log file
func (loggy Logger) Rec(msg string, value interface{}) {
	if value != nil {
		loggy.Printf("%s %v", msg, value)
	} else {
		loggy.Println(msg)
	}
}
