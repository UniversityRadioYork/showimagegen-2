package logging

import "log"

func logWithLevel(level string, content interface{}) {
	log.Printf("[%s] %v\n", level, content)
}

func Debug(content interface{}) {
	logWithLevel("DEBUG", content)
}

func Info(content interface{}) {
	logWithLevel("INFO", content)
}

func Warning(content interface{}) {
	logWithLevel("WARNING", content)
}

func Error(err error) {
	logWithLevel("ERROR", err)
}

func Fatal(err error) {
	log.Fatalf("[FATAL] %v\n", err)
}
