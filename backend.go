package gologs

import (
	"fmt"
	"log"
	"os"
	"time"
)

var output *log.Logger

func InfoBackendf(message string, args ...any) {
	InfoBackend(fmt.Sprintf(message, args...))
}

func InfoBackend(message string) {
	printBackendWithLevel("INFO", message)
}

func WarningBackendf(message string, args ...any) {
	WarningBackend(fmt.Sprintf(message, args...))
}

func WarningBackend(message string) {
	printBackendWithLevel("WARN", message)
}

func ErrorBackendf(message string, args ...any) {
	ErrorBackend(fmt.Sprintf(message, args...))
}

func ErrorBackend(message string) {
	printBackendWithLevel("ERROR", message)
}

func DebugBackendf(message string, args ...any) {
	DebugBackend(fmt.Sprintf(message, args...))
}

func DebugBackend(message string) {
	if os.Getenv(DebugEnvVar) == "true" {
		printBackendWithLevel("DEBUG", message)
	}
}

func FatalBackendf(message string, args ...any) {
	FatalBackend(fmt.Sprintf(message, args...))
}

func FatalBackend(message string) {
	printBackendWithLevel("FATAL", message)
}

func printBackendWithLevel(level, message string) {
	if output == nil {
		output = log.New(
			os.Stdout,
			"",
			0,
		)
	}
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(output.Writer(), "[%s] [%-5s] %s\n",
		timestamp, level, message)
}
