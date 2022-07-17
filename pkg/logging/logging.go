package logging

import "log"

func logWithLevel(level string, content interface{}) {
	log.Printf("[%s] %v\n", level, content)
}

// Debug logs at Debug level.
func Debug(content interface{}) {
	logWithLevel("DEBUG", content)
}

// Info logs at info level.
func Info(content interface{}) {
	logWithLevel("INFO", content)
}

// Warning logs at warning level.
func Warning(content interface{}) {
	logWithLevel("WARNING", content)
}

// Error logs at error level (recoverable).
func Error(err error) {
	logWithLevel("ERROR", err)
}

// Fatal logs at fatal err and panics.
func Fatal(err error) {
	log.Fatalf("[FATAL] %v\n", err)
}
