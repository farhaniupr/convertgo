package timechange

import (
	"strings"
	"time"

	"gitee.com/go-package/carbon"
	"github.com/farhaniupr/convertgo"
)

// using carbon
func PostDateString(target_time string) string {
	result := carbon.Parse(target_time).DiffForHumans()
	return result
}

func InterfaceUTCtoGMT7(target interface{}) string {

	if target == nil {
		return ""
	} else if convertgo.ItString(target) == "" {
		return ""
	} else {

		t := target.(time.Time)

		loc, _ := time.LoadLocation("Asia/Jakarta")

		t = t.In(loc)

		return t.Format("2006-01-02 15:04:02")
	}
}

func InterfaceTimeRemoveWordUtc(target interface{}) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(convertgo.ItString(target), "T", " "), "Z", ""), " +0000 U C", "")
}
