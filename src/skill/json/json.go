package json

import (
	"encoding/json"
	"bytes"
)

// 接口返回失败时，data的数据
var FailDataValue = []byte("[]")

// 接口返回数据
type ApiResult struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data ApiResultData `json:"data"`
}

// 当接口返回成功时，携带的数据
type ApiResultData struct {
	ApiResultDataValue
}

type ApiResultDataValue struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// 实现Json的接口来实现针对不同结果进行不同的解析
func (data *ApiResultData) UnmarshalJSON(value []byte) error {
	// 处理[]的情况
	if bytes.Compare(value, FailDataValue) == 0 {
		return nil
	}

	return json.Unmarshal(value, &data.ApiResultDataValue)
}
