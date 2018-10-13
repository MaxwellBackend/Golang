package _struct

import (
	"testing"
	"fmt"
)

var student = &Student{1, "张三", 20, 100}

func TestStudentSimpleData(t *testing.T) {


	data, err := StudentSimpleData(student)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("simpleData=%s\n", data)
}

func TestStudentFullData(t *testing.T) {
	data, err := StudentFullData(student)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("fullData=%s\n", data)
}