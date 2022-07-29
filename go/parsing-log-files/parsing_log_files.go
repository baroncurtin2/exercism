package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re, _ := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re, _ := regexp.Compile(`(<[~\*=-]*>)`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re, _ := regexp.Compile(`(?i)".*password.*"`)

	count := 0

	for _, l := range lines {
		if re.MatchString(l) {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re, _ := regexp.Compile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re, _ := regexp.Compile(`User\s+(\w+\d+)`)

	for i, l := range lines {
		if re.MatchString(l) {
			matches := re.FindStringSubmatch(l)
			user := matches[1]

			lines[i] = fmt.Sprintf("[USR] %s %s", user, l)
		}
	}
	return lines
}
