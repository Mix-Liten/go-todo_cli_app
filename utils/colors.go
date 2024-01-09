package utils

import "fmt"

// Constants for text color representations
const (
	colorDefault = "\x1b[39m" // Default color

	colorRed   = "\x1b[91m" // Red color
	colorGreen = "\x1b[32m" // Green color
	colorBlue  = "\x1b[94m" // Blue color
	colorGray  = "\x1b[90m" // Gray color
)

// red changes the text color to red
func Red(s string) string {
	return fmt.Sprintf("%s%s%s", colorRed, s, colorDefault)
}

// green changes the text color to green
func Green(s string) string {
	return fmt.Sprintf("%s%s%s", colorGreen, s, colorDefault)
}

// blue changes the text color to blue
func Blue(s string) string {
	return fmt.Sprintf("%s%s%s", colorBlue, s, colorDefault)
}

// gray changes the text color to gray
func Gray(s string) string {
	return fmt.Sprintf("%s%s%s", colorGray, s, colorDefault)
}
