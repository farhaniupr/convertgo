package convertgo_test

import (
	"testing"

	"github.com/farhaniupr/convertgo"
)

// interface{} to string
func ItString(t *testing.T) {
	value_testing := "data"
	var value_testing_2nd *string

	result := convertgo.ItString(value_testing)
	result_testing := convertgo.ItString(value_testing_2nd)

	if result != "data" {
		t.Fatalf("error")
	}

	if result_testing != "" {
		t.Fatalf("error")
	}

}
