package _for

import (
	"testing"
	"fmt"
)

func fillArr() []string {
	var size = 100000
	var arr = make([]string, 0, size)
	for i := 0; i < size; i++ {
		arr = append(arr, fmt.Sprintf("Number_%v", i))
	}

	return arr
}

func BenchmarkForSliceWithIndex(b *testing.B) {
	arr := fillArr()
	for i := 0; i < b.N; i++ {
		ForSliceWithIndex(arr)
	}

}

func BenchmarkForSliceWithRange(b *testing.B) {
	arr := fillArr()
	for i := 0; i < b.N; i++ {
		ForSliceWithRange(arr)
	}
}
