package convertgo_test

import (
	"log"
	"testing"

	"github.com/farhaniupr/convertgo"
	"github.com/farhaniupr/convertgo/thousandconvert"
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

func TestPointStringToString(t *testing.T) {
	var value_star_string *string
	value_testing := "data"
	value_star_string = &value_testing

	result := convertgo.PointStringToString(value_star_string)
	result_testing := convertgo.PointStringToString(nil)

	if result != "data" {
		t.Fatalf("error")
	}

	if result_testing != "" {
		t.Fatalf("error")
	}

}

func TestPointStringToPointString(t *testing.T) {
	var value_star_string *string
	value_testing := "data"
	value_star_string = &value_testing

	result := convertgo.PointStringToPointString(value_star_string, "http", "//")
	result_testing := convertgo.PointStringToPointString(nil)

	var r_result string
	r_result = *result

	if r_result != "//httpdata" {
		t.Fatalf("error")
	}

	if result_testing != nil {
		t.Fatalf("error")
	}
}

func TestThousnad(t *testing.T) {
	log.Println("testing 29090: ", thousandconvert.NearestThousandFormatInterface("29090"))
	log.Println("testing 29090: ", thousandconvert.NearestThousandFormat("29090"))
	log.Println("testing 29090: ", thousandconvert.NearestThousandFormatInteger(29090))

	log.Println("testing 2099: ", thousandconvert.NearestThousandFormatInterface("2099"))
	log.Println("testing 2099: ", thousandconvert.NearestThousandFormat("2099"))
	log.Println("testing 2099: ", thousandconvert.NearestThousandFormatInteger(2099))
}
