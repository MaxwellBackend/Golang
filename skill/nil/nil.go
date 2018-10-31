package nil

import (
	"reflect"
)

// 常用的Nil判断
func NormalNilJudge(value interface{}) bool {
	return value == nil
}

// 严格的Nil判断
func StrictNilJudge(value interface{}) bool {
	return value == nil || reflect.ValueOf(value).IsNil()
}
