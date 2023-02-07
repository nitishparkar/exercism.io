package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	matched, _ := regexp.MatchString(`^\[TRC\]|^\[DBG\]|^\[INF\]|^\[WRN\]|^\[ERR\]|^\[FTL\]`, text)
	return matched
}

func SplitLogLine(text string) []string {
	return regexp.MustCompile(`<(\~|\*|\=|\-)*>`).Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0

	r, _ := regexp.Compile(`(?i)".*password.*"`)
	for _, line := range lines {
		if matched := r.MatchString(line); matched {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	r := regexp.MustCompile(`end-of-line[\d]+`)
	return r.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	processedLines := make([]string, len(lines))

	r := regexp.MustCompile(`User\s+([^\s]+)`)
	for i, line := range lines {
		if submatches := r.FindStringSubmatch(line); len(submatches) > 0 {
			processedLines[i] = "[USR] " + submatches[1] + " " + line
		} else {
			processedLines[i] = line
		}
	}

	return processedLines
}
