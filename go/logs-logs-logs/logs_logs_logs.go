package logs

import "strings"

// Application identifies the application emitting the given log.
func Application(log string) string {
	index := strings.IndexAny(log, string([]rune{'\U00002757', '\U0001F50D', '\U00002600'}))
	if index == -1 {
		return "default"
	}

	switch []rune(log)[index] {
	case '\U00002757':
		return "recommendation"
	case '\U0001F50D':
		return "search"
	case '\U00002600':
		return "weather"
	default:
		return "default"
	}
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return len([]rune(log)) <= limit
}
