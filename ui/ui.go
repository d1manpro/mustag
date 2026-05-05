package ui

import "github.com/fatih/color"

var (
	Bold   = color.New(color.Bold).SprintFunc()
	Italic = color.New(color.Italic).SprintFunc()

	Red    = color.New(color.FgRed).SprintFunc()
	Green  = color.New(color.FgGreen).SprintFunc()
	Blue   = color.New(color.FgBlue).SprintFunc()
	Yellow = color.New(color.FgYellow).SprintFunc()
)
