package _struct

import "encoding/json"

// 学生信息
type Student struct {
	Id    uint32
	Name  string
	Age   uint32
	Score uint32
}

// 学生简单信息
func StudentSimpleData(student *Student) ([]byte, error) {
	var studentSimple struct {
		Id   uint32
		Name string
	}

	studentSimple.Id = student.Id
	studentSimple.Name = student.Name

	return json.Marshal(studentSimple)
}

// 学生全量信息
func StudentFullData(student *Student) ([]byte, error) {
	return json.Marshal(student)
}
