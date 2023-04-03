package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	r := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return r.MatchString(text)
}

func SplitLogLine(text string) []string {
	r := regexp.MustCompile(`<[-=~*]*>`)

	return strings.Split(r.ReplaceAllString(text, "|"), "|")
}

func CountQuotedPasswords(lines []string) int {
	r := regexp.MustCompile(`"[^"]*password[^"]*"`)

	count := 0

	for _, line := range lines {
		line = strings.ToLower(line)
		if r.MatchString(line) {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	r := regexp.MustCompile(`end-of-line[0-9]*`)

	return r.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)
	ret := []string{}
	for _, line := range lines {
		founds := re.FindStringSubmatch(line)
		if founds != nil {
			line = fmt.Sprintf("[USR] %s %s", founds[1], line)
		}
		ret = append(ret, line)
	}
	return ret
}
