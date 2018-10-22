package _for

// 用索引遍历数组
func ForSliceWithIndex(arr []string) {
	for i := 0; i < len(arr); i++ {
		_, _ = i, arr[i]
	}
}

// 用Range遍历索引
func ForSliceWithRange(arr []string) {
	for i, v := range arr {
		_, _ = i, v
	}
}
