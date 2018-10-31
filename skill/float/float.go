package float

import (
	"github.com/shopspring/decimal"
)

func FloatMulti(value float64) int {
	return int(value*100)
}

func FloatDecimalMulti(value float64) int {
	decimalValue := decimal.NewFromFloat(value)
	decimal100 := decimal.NewFromFloat(100)
	result, _ := decimalValue.Mul(decimal100).Float64()

	return int(result)
}