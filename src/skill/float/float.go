package float

import "math"

func FloatMulti(value float64) float64 {
	return value*100
}

func FloatCeilMulti(value float64) int {
	return int(math.Ceil(value*100))
}