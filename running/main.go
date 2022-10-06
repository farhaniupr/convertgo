package main

import (
	"log"
	"time"

	"github.com/farhaniupr/convertgo/thousandconvert"
	"github.com/farhaniupr/convertgo/timechange"
)

func main() {

	// example

	log.Println(timechange.InterfaceUTCtoGMT7(nil))

	log.Println(timechange.InterfaceUTCtoGMT7(time.Now()))

	log.Println(thousandconvert.NearestThousandFormat("2080"))

	log.Println(thousandconvert.NearestThousandFormatInteger(2080))

	log.Println(thousandconvert.NearestThousandFormatInterface(2180))

	log.Println(thousandconvert.NearestThousandFormatString(2080))
}
