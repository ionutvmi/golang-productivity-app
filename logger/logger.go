package logger

import (
	"io"
	"log"
	"os"
)

var appLogFile *os.File

func Initialize() {
	f, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	writer := io.MultiWriter(f)
	log.SetOutput(writer)
	log.SetPrefix("APP: ")

	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)

	appLogFile = f
}

// Close is intended to be called from main
func Close() {
	appLogFile.Close()
}
