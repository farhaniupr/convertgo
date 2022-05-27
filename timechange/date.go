package timechange

import "github.com/golang-module/carbon"

// using carbon
func PostDateString(target_time string) string {
	result := carbon.Parse(target_time).DiffForHumans()
	return result
}
