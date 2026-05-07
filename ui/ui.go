package ui

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

var (
	keyColor    = color.New(color.FgGreen, color.Bold)
	valueColor  = color.New(color.FgWhite)
	headerColor = color.New(color.FgHiCyan, color.Bold)
	warnColor   = color.New(color.FgYellow, color.Bold)
	errColor    = color.New(color.FgRed, color.Bold)
)

func KeyValue(key, value string) {
	if value == "" {
		return
	}

	fmt.Fprintf(os.Stdout, "%s:\t%s\n",
		keyColor.Sprint(key),
		valueColor.Sprint(value),
	)
}

func Raw(value string) {
	fmt.Fprintln(os.Stdout, value)
}

func Info(msg string) {
	fmt.Fprintln(os.Stdout, msg)
}

func Warn(msg string) {
	fmt.Fprintln(os.Stderr, warnColor.Sprintf("WARN: %s", msg))
}

func Error(msg string) {
	fmt.Fprintln(os.Stderr, errColor.Sprintf("ERROR: %s", msg))
}

func Header(msg string) {
	fmt.Fprintln(os.Stdout, headerColor.Sprintf("=== %s ===", msg))
}

func NewLine() {
	fmt.Fprintln(os.Stdout)
}
