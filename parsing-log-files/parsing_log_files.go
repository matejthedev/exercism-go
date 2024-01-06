package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<(?:[~=,*-]*)>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	re := regexp.MustCompile(`(?i)"[^"]*password[^"]*"`)
	for _, v := range lines {
		str := re.FindAllString(v, -1)
		count += len(str)
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*(\s*)`)
	return re.ReplaceAllString(text, "$1")
}

func TagWithUserName(lines []string) []string {
	regex := regexp.MustCompile(`User\s+(\w*)`)
	newLines := make([]string, len(lines))
	copy(newLines, lines)
	for i, line := range newLines {
		if submatches := regex.FindStringSubmatch(line); submatches != nil {
			newLines[i] = fmt.Sprintf("[USR] %s %s", submatches[1], line)
		}
	}
	return newLines
}
