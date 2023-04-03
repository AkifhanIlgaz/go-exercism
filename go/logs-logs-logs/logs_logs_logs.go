package logs

import "strings"

// Application identifies the application emitting the given log.
func Application(log string) string {

	contains := func(log rune) bool {
		emits := []rune{'\u2757', '🔍', '\u2600'}

		for _, emit := range emits {
			if emit == log {
				return true
			}
		}

		return false
	}

	for _, r := range log {
		if contains(r) {
			var app string

			switch r {
			case '\u2757':
				app = "recommendation"
			case '🔍':
				app = "search"
			case '\u2600':
				app = "weather"
			}

			return app
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return len(strings.Split(log, "")) <= limit
}
