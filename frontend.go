package gologs

import (
	"fmt"
	"os"
)

func Infof(message string, args ...any) {
	Info(fmt.Sprintf(message, args...))
}

func Info(message string) {
	printWithLevel("INFO", message, "\033[34m")
}

func Warningf(message string, args ...any) {
	Warning(fmt.Sprintf(message, args...))
}

func Warning(message string) {
	printWithLevel("WARNING", message, "\033[33m")
}

func Errorf(message string, args ...any) {
	Error(fmt.Sprintf(message, args...))
}

func Error(message string) {
	printWithLevel("ERROR", message, "\033[31m")
}

func Debugf(message string, args ...any) {
	Debug(fmt.Sprintf(message, args...))
}

func Debug(message string) {
	if os.Getenv("DEBUG") == "true" {
		printWithLevel("DEBUG", message, "\033[36m")
	}
}

func Fatalf(message string, args ...any) {
	printWithLevel("FATAL", fmt.Sprintf(message, args...), "\033[35m")
	os.Exit(1)
}

func Fatal(message string) {
	printWithLevel("FATAL", message, "\033[35m")
	os.Exit(1)
}

func printWithLevel(level, message, color string) {
	fmt.Printf("%s%-7s => %s\033[0m\n", color, level, message)
}
