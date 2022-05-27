package casestring

import "strings"

// lower value string from type inteface{} to lower
func Lower(v interface{}) interface{} {
	switch v := v.(type) {
	case []interface{}:
		for i := range v {
			v[i] = Lower(v[i])
		}
		return v
	case map[string]interface{}:
		lv := make(map[string]interface{}, len(v))
		for mk, mv := range v {
			lv[strings.ToLower(mk)] = mv
		}
		return lv
	default:
		return v
	}
}

// upcase value string from type inteface{} to upcase
func Upcase(v interface{}) interface{} {
	switch v := v.(type) {
	case []interface{}:
		for i := range v {
			v[i] = Upcase(v[i])
		}
		return v
	case map[string]interface{}:
		lv := make(map[string]interface{}, len(v))
		for mk, mv := range v {
			lv[strings.ToUpper(mk)] = mv
		}
		return lv
	default:
		return v
	}
}
