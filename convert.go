package convertgo

import "fmt"

// interface{} to string
func ItString(value interface{}) string {
	if value == nil {
		return ""
	} else {
		data := fmt.Sprintf("%v", value)
		return data
	}

}

// interface{} to interface{}(integer)
func ItInt(i interface{}) interface{} {
	if i == nil {
		return 0
	} else {

		return i
	}
}

// *string to string
func PointStringToString(value *string) string {
	var value_return string
	if value == nil {
		value_return = ""
	} else {
		valuestar := *value
		value_return = valuestar
	}
	return value_return
}

// *string to *string
func PointStringToPointString(value *string, insert ...string) *string {
	var value_return *string
	if value == nil {
		value_return = nil
	} else if len(insert) > 0 {
		for i := 0; i < len(insert); i++ {
			if i > 0 {
				value_next := *value_return
				valuenotnull := insert[i] + value_next
				value_return = &valuenotnull
			} else {
				valuestar := *value
				valuenotnull := insert[i] + valuestar
				value_return = &valuenotnull
			}
		}
	} else if len(insert) == 0 {
		valuestar := *value
		valuenotnull := valuestar
		value_return = &valuenotnull
	}
	return value_return
}
