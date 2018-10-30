package exists

import "testing"

func BenchmarkExistsWithInterface(b *testing.B) {
	var set = make(map[uint32]interface{})
	for i:=uint32(1);i<uint32(b.N);i++ {
		set[i] = nil
	}

	for i:=0;i<b.N;i++ {
		ExistsWithInterface(set, uint32(i))
	}
}

func BenchmarkExistsWithStruct(b *testing.B) {
	var set = make(map[uint32]struct{})
	for i:=uint32(1);i<uint32(b.N);i++ {
		set[i] = struct{}{}
	}

	for i:=0;i<b.N;i++ {
		ExistsWithStruct(set, uint32(i))
	}
}

func BenchmarkExistsWithBool(b *testing.B) {
	var set = make(map[uint32]bool)
	for i:=uint32(1);i<uint32(b.N);i++ {
		set[i] = true
	}

	for i:=0;i<b.N;i++ {
		ExistsWithBool(set, uint32(i))
	}
}

func BenchmarkExistsWithInt(b *testing.B) {
	var set = make(map[uint32]int)
	for i:=uint32(1);i<uint32(b.N);i++ {
		set[i] = 1
	}

	for i:=0;i<b.N;i++ {
		ExistsWithInt(set, uint32(i))
	}
}
