package iocli

import (
	"fmt"

	"github.com/fatih/color"
)

var fail = "x"
var ambiguous = "-"
var pass = "*"

var infoColor = color.New(color.FgBlue).SprintFunc()
var successColor = color.New(color.FgGreen).SprintFunc()
var errorColor = color.New(color.FgRed).SprintFunc()
var promptColor = color.New(color.FgMagenta).SprintFunc()

func printOut(prefix string, format string, a ...interface{}) {
	fmt.Printf("%s %s", prefix, fmt.Sprintf(format, a...))
}

// Success prints info to the console
func Success(format string, a ...interface{}) {
	printOut(successColor("*>"), fmt.Sprintf("%s\n", format), a...)
}

// Info prints info to the console
func Info(format string, a ...interface{}) {
	printOut(infoColor("->"), fmt.Sprintf("%s\n", format), a...)
}

// Error prints an error message to the concole
func Error(format string, a ...interface{}) {
	printOut(errorColor("!>"), fmt.Sprintf("%s\n", format), a...)
}

// Errors prints a list of errors
func Errors(errs []string) {
	for _, v := range errs {
		Error("  %s", v)
	}
}

func printPrompt(format string, a ...interface{}) {
	printOut(promptColor("?>"), fmt.Sprintf("%s: ", format), a...)
}
