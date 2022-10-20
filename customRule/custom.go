package customRule

import (
	"System/System"
	"errors"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"strings"
)

func Init() {
	govalidator.AddCustomRule("unique", func(field string, rule string, message string, value interface{}) error {
		tableName := strings.TrimPrefix(rule, "unique:") //handle other error
		var r int64
		System.DB.Table(tableName).Where(field+" = ?", value).Count(&r)
		if r != 0 {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("the %s field must be unique", field)
		}
		return nil
	})
}
