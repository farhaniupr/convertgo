package convertgo_test

import (
	"log"
	"testing"

	"github.com/farhaniupr/convertgo"
)

// interface{} to string
func TestItString(t *testing.T) {
	value_testing := "data"

	result := convertgo.ItString(value_testing)
	result_testing := convertgo.ItString(nil)

	log.Println(result_testing)

	if result != "data" {
		t.Fatalf("error")
	}

	if result_testing != "" {
		t.Fatalf("error")
	}

}
