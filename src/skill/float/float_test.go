package float

import (
	"testing"
	"fmt"
)

func TestFloatMulti(t *testing.T) {
	fmt.Printf("19.99 * 100 = %v\n", FloatMulti(19.99))
	fmt.Printf("69.99 * 100 = %v\n", FloatMulti(69.99))
}

func TestFloatCeilMulti(t *testing.T) {
	fmt.Printf("19.99 * 100 = %v\n", FloatCeilMulti(19.99))
	fmt.Printf("69.99 * 100 = %v\n", FloatCeilMulti(69.99))
}
