package exists

// 判断是否存在，用interface做为值
func ExistsWithInterface(set map[uint32]interface{}, key uint32) bool {
	_, found := set[key]
	return found
}

// 判断是否存在，用struct做为值
func ExistsWithStruct(set map[uint32]struct{}, key uint32) bool {
	_, found := set[key]
	return found
}

// 判断是否存在，用bool做为值
func ExistsWithBool(set map[uint32]bool, key uint32) bool {
	_, found := set[key]
	return found
}

// 判断是否存在，用int做为值
func ExistsWithInt(set map[uint32]int, key uint32) bool {
	_, found := set[key]
	return found
}
