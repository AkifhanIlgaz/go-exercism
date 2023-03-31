package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	customer = strings.ToUpper(customer)
	return fmt.Sprintf("Welcome to the Tech Palace, %s", customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var sb strings.Builder

	stars := fmt.Sprint(strings.Repeat("*", numStarsPerLine))

	sb.WriteString(stars + "\n")
	sb.WriteString(welcomeMsg + "\n")
	sb.WriteString(stars)

	return sb.String()
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	cleanedMessage := strings.ReplaceAll(oldMsg, "*", "")
	cleanedMessage = strings.TrimSpace(cleanedMessage)
	return cleanedMessage
}
