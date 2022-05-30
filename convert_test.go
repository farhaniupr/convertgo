package convertgo_test

import (
	"testing"

	"github.com/farhaniupr/convertgo"
)

// go test -v

func TestItString(t *testing.T) {
	value_testing := "data"

	result := convertgo.ItString(value_testing)
	result_testing := convertgo.ItString(nil)

	if result != "data" {
		t.Fatalf("error")
	}

	if result_testing != "" {
		t.Fatalf("error")
	}

}
