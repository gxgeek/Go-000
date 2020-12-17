package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

func BindAndVailValue(c *gin.Context, form interface{}) error {
	err := c.Bind(form)
	if err != nil {
		return err
	}
	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return err
	}

	if !check {
		MarkErrors(valid.Errors)
		return err
	}

	return nil
}
