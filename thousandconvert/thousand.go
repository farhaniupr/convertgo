package thousandconvert

import (
	"math"
	"strconv"
	"strings"

	"github.com/farhaniupr/convertgo"
)

func RoundPrec(x float64, prec int) float64 {
	if math.IsNaN(x) || math.IsInf(x, 0) {
		return x
	}

	sign := 1.0
	if x < 0 {
		sign = -1
		x *= -1
	}

	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)

	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow * sign
}

// format thousand string to string
func NumberFormat(number float64, decimals int, decPoint, thousandsSep string) string {
	if math.IsNaN(number) || math.IsInf(number, 0) {
		number = 0
	}

	var ret string
	var negative bool

	if number < 0 {
		number *= -1
		negative = true
	}

	d, fract := math.Modf(number)

	if decimals <= 0 {
		fract = 0
	} else {
		pow := math.Pow(10, float64(decimals))
		fract = RoundPrec(fract*pow, 0)
	}

	if thousandsSep == "" {
		ret = strconv.FormatFloat(d, 'f', 0, 64)
	} else if d >= 1 {
		var x float64
		for d >= 1 {
			d, x = math.Modf(d / 1000)
			x = x * 1000
			ret = strconv.FormatFloat(x, 'f', 0, 64) + ret
			if d >= 1 {
				ret = thousandsSep + ret
			}
		}
	} else {
		ret = "0"
	}

	fracts := strconv.FormatFloat(fract, 'f', 0, 64)

	// "0" pad left
	for i := len(fracts); i < decimals; i++ {
		fracts = "0" + fracts
	}

	ret += decPoint + fracts

	if negative {
		ret = "-" + ret
	}
	return ret
}

func RoundInt(input float64) int {
	var result float64

	if input < 0 {
		result = math.Ceil(input - 0.5)
	} else {
		result = math.Floor(input + 0.5)
	}

	// only interested in integer, ignore fractional
	i, _ := math.Modf(result)

	return int(i)
}

func FormatNumber(input float64) string {
	x := RoundInt(input)
	xFormatted := NumberFormat(float64(x), 2, ".", ",")
	return xFormatted
}

func NearestThousandFormat(num string) string {

	flo, _ := strconv.ParseFloat(num, 64)

	if math.Abs(flo) < 999.5 {
		xNum := FormatNumber(flo)
		xNumStr := xNum[:len(xNum)-3]
		return string(xNumStr)
	}

	xNum := FormatNumber(flo)
	// first, remove the .00 then convert to slice
	xNumStr := xNum[:len(xNum)-3]
	xNumCleaned := strings.Replace(xNumStr, ",", " ", -1)
	xNumSlice := strings.Fields(xNumCleaned)
	count := len(xNumSlice) - 2
	unit := [4]string{"k", "m", "b", "t"}
	xPart := unit[count]

	afterDecimal := ""
	if xNumSlice[1][0] != 0 {
		afterDecimal = "." + string(xNumSlice[1][0])
	}
	final := xNumSlice[0] + afterDecimal + xPart

	if len(num) == 4 {
		if convertgo.ItInt(string(num[1])) == 0 && convertgo.ItInt(string(num[2])) > 0 {
			final = string(num[0]) + xPart
		}
	} else if len(num) == 5 {
		if convertgo.ItInt(string(num[2])) == 0 && convertgo.ItInt(string(num[3])) > 0 {
			final = string(num[0]) + string(num[1]) + xPart
		}
	} else if len(num) == 6 {
		if convertgo.ItInt(string(num[3])) == 0 && convertgo.ItInt(string(num[4])) > 0 {
			final = string(num[0]) + string(num[1]) + string(num[2]) + xPart
		}
	}

	return final
}

// format thousand interface{} to interface{}
func NearestThousandFormatInterface(num interface{}) interface{} {

	if num != nil {

		datainterface := convertgo.ItString(num)

		flo, _ := strconv.ParseFloat(datainterface, 64)

		if math.Abs(flo) < 999.5 {
			xNum := FormatNumber(flo)
			xNumStr := xNum[:len(xNum)-3]
			return string(xNumStr)
		}

		xNum := FormatNumber(flo)
		// first, remove the .00 then convert to slice
		xNumStr := xNum[:len(xNum)-3]
		xNumCleaned := strings.Replace(xNumStr, ",", " ", -1)
		xNumSlice := strings.Fields(xNumCleaned)
		count := len(xNumSlice) - 2
		unit := [4]string{"k", "m", "b", "t"}
		xPart := unit[count]

		afterDecimal := ""
		if xNumSlice[1][0] != 0 {
			afterDecimal = "." + string(xNumSlice[1][0])
		}
		final := xNumSlice[0] + afterDecimal + xPart

		if len(datainterface) == 4 {
			if convertgo.ItInt(string(datainterface[1])) == 0 && convertgo.ItInt(string(datainterface[2])) > 0 {
				final = string(datainterface[0]) + xPart
			}
		} else if len(datainterface) == 5 {
			if convertgo.ItInt(string(datainterface[2])) == 0 && convertgo.ItInt(string(datainterface[3])) > 0 {
				final = string(datainterface[0]) + string(datainterface[1]) + xPart
			}
		} else if len(datainterface) == 6 {
			if convertgo.ItInt(string(datainterface[3])) == 0 && convertgo.ItInt(string(datainterface[4])) > 0 {
				final = string(datainterface[0]) + string(datainterface[1]) + string(datainterface[2]) + xPart
			}
		}

		return final
	} else {
		return ""
	}
}

func NearestThousandFormatString(num interface{}) string {

	if num != nil {

		datainterface := convertgo.ItString(num)

		flo, _ := strconv.ParseFloat(datainterface, 64)

		if math.Abs(flo) < 999.5 {
			xNum := FormatNumber(flo)
			xNumStr := xNum[:len(xNum)-3]
			return string(xNumStr)
		}

		xNum := FormatNumber(flo)
		// first, remove the .00 then convert to slice
		xNumStr := xNum[:len(xNum)-3]
		xNumCleaned := strings.Replace(xNumStr, ",", " ", -1)
		xNumSlice := strings.Fields(xNumCleaned)
		count := len(xNumSlice) - 2
		unit := [4]string{"k", "m", "b", "t"}
		xPart := unit[count]

		afterDecimal := ""
		if xNumSlice[1][0] != 0 {
			afterDecimal = "." + string(xNumSlice[1][0])
		}
		final := xNumSlice[0] + afterDecimal + xPart

		if len(datainterface) == 4 {
			if convertgo.ItInt(string(datainterface[1])) == 0 && convertgo.ItInt(string(datainterface[2])) > 0 {
				final = string(datainterface[0]) + xPart
			}
		} else if len(datainterface) == 5 {
			if convertgo.ItInt(string(datainterface[2])) == 0 && convertgo.ItInt(string(datainterface[3])) > 0 {
				final = string(datainterface[0]) + string(datainterface[1]) + xPart
			}
		} else if len(datainterface) == 6 {
			if convertgo.ItInt(string(datainterface[3])) == 0 && convertgo.ItInt(string(datainterface[4])) > 0 {
				final = string(datainterface[0]) + string(datainterface[1]) + string(datainterface[2]) + xPart
			}
		}

		return final
	} else {
		return ""
	}

}

// format thousand int to string
func NearestThousandFormatInteger(num int) string {

	datainterface := convertgo.ItString(num)

	flo := float64(num)

	if math.Abs(flo) < 999.5 {
		xNum := FormatNumber(flo)
		xNumStr := xNum[:len(xNum)-3]
		return string(xNumStr)
	}

	xNum := FormatNumber(flo)
	// first, remove the .00 then convert to slice
	xNumStr := xNum[:len(xNum)-3]
	xNumCleaned := strings.Replace(xNumStr, ",", " ", -1)
	xNumSlice := strings.Fields(xNumCleaned)
	count := len(xNumSlice) - 2
	unit := [4]string{"k", "m", "b", "t"}
	xPart := unit[count]

	afterDecimal := ""
	if xNumSlice[1][0] != 0 {
		afterDecimal = "." + string(xNumSlice[1][0])
	}
	final := xNumSlice[0] + afterDecimal + xPart

	if len(datainterface) == 4 {
		if convertgo.ItInt(string(datainterface[1])) == 0 && convertgo.ItInt(string(datainterface[2])) > 0 {
			final = string(datainterface[0]) + xPart
		}
	} else if len(datainterface) == 5 {
		if convertgo.ItInt(string(datainterface[2])) == 0 && convertgo.ItInt(string(datainterface[3])) > 0 {
			final = string(datainterface[0]) + string(datainterface[1]) + xPart
		}
	} else if len(datainterface) == 6 {
		if convertgo.ItInt(string(datainterface[3])) == 0 && convertgo.ItInt(string(datainterface[4])) > 0 {
			final = string(datainterface[0]) + string(datainterface[1]) + string(datainterface[2]) + xPart
		}
	}

	return final
}
