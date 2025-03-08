package utils

import "log"

func Log(message string, file string, line int) {
	log.Printf("%s:%d: %s", file, line, message)
}

func LogError(err error, file string, line int) {
	log.Printf("%s:%d: %v", file, line, err)
}
