package formatmoney

import (
	"strings"

	"github.com/dustin/go-humanize"
)

// format rupiah using go-humanize
func FormatRupiah(amount float64) string {
	humanizeValue := humanize.CommafWithDigits(amount, 0)
	stringValue := strings.Replace(humanizeValue, ",", ".", -1)
	return stringValue
}
