package timechange

import (
	"time"

	"gitee.com/go-package/carbon"
)

// using carbon
func PostDateString(target_time string) string {
	result := carbon.Parse(target_time).DiffForHumans()
	return result
}

func InterfaceUTCtoGMT7(target interface{}) string {

	t := target.(time.Time)

	loc, _ := time.LoadLocation("Asia/Jakarta")

	t = t.In(loc)

	return t.Format("2006-01-02 15:04:02")
}
