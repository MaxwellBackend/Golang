package json

import (
	"testing"
	"encoding/json"
)

// 测试成功返回值的解析
func Test_ResultOk(t *testing.T) {
	data := []byte(`{
		"code": 0,
		"msg": "ok",
		"data": {
			"id": 1,
			"name": "test"
		}
	}`)

	result := &ApiResult{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Error(err)
	}
}

// 测试失败返回值的解析
func Test_ResultFail(t *testing.T) {
	data := []byte(`{
		"code": 1,
		"msg": "fail",
		"data": []
	}`)

	result := &ApiResult{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Error(err)
	}
}
