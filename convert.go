package convertgo

import "fmt"

func ItString(value interface{}) string {
	if value == nil {
		return ""
	} else {
		data := fmt.Sprintf("%v", value)
		return data
	}

}

func ItInt(i interface{}) interface{} {

	if i == nil {
		return 0
	} else {

		return i
	}
}
