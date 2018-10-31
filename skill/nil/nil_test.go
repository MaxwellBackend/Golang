package nil

import (
	"testing"
	"bytes"
	"fmt"
)

var value *bytes.Buffer

// 常用的Nil判断
func TestNormalNilJudge(t *testing.T)  {
	// 直接传Nil
	if NormalNilJudge(nil) != true {
		t.Errorf("value1 != nil")
	}

	// 间接传入Nil
	result := NormalNilJudge(value)

	fmt.Printf("NormalNilJudge (value == nil) = %v\n", result)

	if NormalNilJudge(value) == true {
		t.Errorf("value == nil")
	}
}

// 严格的Nil判断
func TestStrictNilJudge(t *testing.T) {
	// 直接传Nil
	if StrictNilJudge(nil) != true {
		t.Errorf("value1 != nil")
	}

	// 间接传入Nil
	result := StrictNilJudge(value)

	fmt.Printf("StrictNilJudge (value == nil) = %v\n", result)

	if result != true {
		t.Errorf("value == nil")
	}
}
