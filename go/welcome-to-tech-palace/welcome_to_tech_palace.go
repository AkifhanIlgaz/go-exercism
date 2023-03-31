package techpalace

import (
	"strings"
)

var sb strings.Builder

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	sb.WriteString("Welcome to the Tech Palace, ")
	sb.WriteString(strings.ToUpper(customer))

	defer sb.Reset()
	return sb.String()
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {

	defer sb.Reset()
	starBorder := strings.Repeat("*", numStarsPerLine)

	sb.WriteString(starBorder)
	sb.WriteByte('\n')
	sb.WriteString(welcomeMsg)
	sb.WriteByte('\n')
	sb.WriteString(starBorder)

	return sb.String()
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	cleanedMessage := strings.ReplaceAll(oldMsg, "*", "")
	cleanedMessage = strings.TrimSpace(cleanedMessage)
	return cleanedMessage
}
