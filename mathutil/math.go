package mathutil

import (
	"fmt"
	"strconv"
)

func Decimal(source float64) float64 {
	str := fmt.Sprintf("%.2f", source)
	res, _ := strconv.ParseFloat(str, 64)
	return res
}
