package extractstring

import "strings"

// https://www.php2golang.com/method/function.explode.html
func Explode(delimiter, text string) []string {
	if len(delimiter) > len(text) {
		return strings.Split(delimiter, text)
	} else {
		return strings.Split(text, delimiter)
	}
}
