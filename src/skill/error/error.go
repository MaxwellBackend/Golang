package error

import (
	"errors"
	"fmt"
	"strings"
)

// 数据库处理
func DatabaseOp(sql string) error {
	if strings.Contains(sql, "admin") {
		return errors.New("database op is fail, reason: XXX")
	}

	return nil
}

type ModelError struct {
	table string
	err   error
}

func (e *ModelError) Error() string {
	return fmt.Sprintf("model op fail, table:%v, err:%v", e.table, e.err)
}

// 模型处理
func ModelOp(table string) error {
	if err := DatabaseOp("select * from " + table); err != nil {
		return &ModelError{table: table, err: err}
	}

	return nil
}