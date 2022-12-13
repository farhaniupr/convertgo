package extractstring

import (
	"regexp"
	"strings"
)

// https://www.php2golang.com/method/function.explode.html
func Explode(delimiter, text string) []string {
	if len(delimiter) > len(text) {
		return strings.Split(delimiter, text)
	} else {
		return strings.Split(text, delimiter)
	}
}

func Strip_tags(content string) string {
	re := regexp.MustCompile(`/\s+/`)
	return re.ReplaceAllString(content, "")
}

func RemoveSpecialCharacter(content string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	processedString := reg.ReplaceAllString(content, "-")
	return processedString
}

func Between(value string, a string, b string) string {
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}

func Before(value string, a string) string {
	pos := strings.Index(value, a)
	if pos == -1 {
		return ""
	}
	return value[0:pos]
}

func After(value string, a string) string {
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:len(value)]
}
