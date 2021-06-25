package logger

import (
	"log"
	"os"
	"strings"
)

var isDebug = strings.ToLower(os.Getenv("DEBUG")) == "true"

func Debug(args ...interface{}) {
	if isDebug {
		log.Printf("DEBUG %+v", args...)
	}
}
