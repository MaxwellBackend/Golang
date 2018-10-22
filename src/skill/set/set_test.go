package set

import "testing"

func BenchmarkIssetWithInterface(b *testing.B) {
	var set = make(map[uint32]interface{})
	for i:=uint32(1);i<uint32(b.N);i++ {
		set[i] = nil
	}

	for i:=0;i<b.N;i++ {
		IssetWithInterface(set, uint32(i))
	}
}

func BenchmarkIssetWithStruct(b *testing.B) {
	var set = make(map[uint32]struct{})
	for i:=uint32(1);i<uint32(b.N);i++ {
		set[i] = struct{}{}
	}

	for i:=0;i<b.N;i++ {
		IssetWithStruct(set, uint32(i))
	}
}

func BenchmarkIssetWithBool(b *testing.B) {
	var set = make(map[uint32]bool)
	for i:=uint32(1);i<uint32(b.N);i++ {
		set[i] = true
	}

	for i:=0;i<b.N;i++ {
		IssetWithBool(set, uint32(i))
	}
}

func BenchmarkIssetWithInt(b *testing.B) {
	var set = make(map[uint32]int)
	for i:=uint32(1);i<uint32(b.N);i++ {
		set[i] = 1
	}

	for i:=0;i<b.N;i++ {
		IssetWithInt(set, uint32(i))
	}
}
