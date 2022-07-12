package logger

import (
	"log"

	"github.com/logrusorgru/aurora"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("Fatal\n%s: %s", aurora.BrightRed(msg), aurora.BrightRed(err))
	}
}

func LogError(err error, msg string) {
	log.Printf("%s: %s", aurora.BrightRed(msg), aurora.BrightRed(err))
}

func LogFatal(msg string) {
	log.Fatal(aurora.BrightRed(msg))
}

func LogOutput(msg string) {
	log.Println(aurora.BrightCyan(msg))
}

func LogInfo(msg string) {
	log.Println(aurora.BrightBlue(msg))
}

// Log logs msg with default color.
func Log(msg string) {
	log.Println(msg)
}

// Success logs msg with green color.
func Success(msg string) {
	log.Println(aurora.Green(msg))
}

// Info logs msg with blue color.
func Info(msg string) {
	log.Println(aurora.Cyan(msg))
}

// Error logs msg with Red color.
func Error(err error, msg string) {
	log.Printf("%s: %s", aurora.Red(msg), aurora.Red(err))
}
