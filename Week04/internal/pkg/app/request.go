package app

import (
	"Go-000/Week04/internal/pkg/logging"
	"github.com/astaxie/beego/validation"
)

func MarkErrors(errors []*validation.Error)  {
	for _,err := range errors {
		logging.InfoF(err.Key, err.Message)
	}
}

