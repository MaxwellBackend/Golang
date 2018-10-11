package error

import (
	"testing"
	"fmt"
)

// 测试没有错误的情况
func TestModelOpSuccess(t *testing.T) {
	if err := ModelOp("user"); err != nil {
		t.Error(err)
	}
}

// 测试有错误的情况
func TestModelOpError(t *testing.T) {
	if err := ModelOp("admin"); err == nil {
		t.Error(err)
	} else {
		fmt.Printf("Err: %v\n", err)
	}
}