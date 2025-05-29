package logger

import "log"

func Info(format string, args ...interface{}) {
	log.Printf("[INFO] "+format, args...)
}

func Error(format string, args ...interface{}) {
	log.Printf("[ERROR] "+format, args...)
}

func Debug(format string, args ...interface{}) {
	log.Printf("[DEBUG] "+format, args...)
}

func Warn(format string, args ...interface{}) {
	log.Printf("[WARN] "+format, args...)
}

func Fatal(format string, args ...interface{}) {
	log.Printf("[FATAL] "+format, args...)
}
