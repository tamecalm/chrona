package ui

import "github.com/fatih/color"

var (
	titleColor   = color.New(color.FgHiCyan, color.Bold)
	successColor = color.New(color.FgHiGreen, color.Bold)
	errorColor   = color.New(color.FgHiRed, color.Bold)
	infoColor    = color.New(color.FgHiYellow)
	labelColor   = color.New(color.FgHiWhite, color.Bold)
	valueColor   = color.New(color.FgHiCyan, color.Bold)
)

func Title(text string) string {
	return titleColor.Sprint(text)
}

func Success(text string) string {
	return successColor.Sprint(text)
}

func Error(text string) string {
	return errorColor.Sprint(text)
}

func Info(text string) string {
	return infoColor.Sprint(text)
}

func Label(text string) string {
	return labelColor.Sprint(text)
}

func Value(text string) string {
	return valueColor.Sprint(text)
}
