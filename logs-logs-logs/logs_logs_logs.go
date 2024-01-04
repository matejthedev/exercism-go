package logs

var mapping = map[rune]string{
	'❗': "recommendation",
	'🔍': "search",
	'☀': "weather",
}

func Application(log string) string {
	for _, v := range log {
		log, ok := mapping[v]
		if ok {
			return log
		}
	}
	return "default"
}

func Replace(log string, oldRune, newRune rune) string {
	newLog := ""
	for _, v := range log {
		if v == oldRune {
			newLog += string(newRune)
		} else {
			newLog += string(v)
		}
	}
	return newLog
}

func WithinLimit(log string, limit int) bool {
	return len([]rune(log)) <= limit
}
