package app

import (
	"Go-000/Week04/internal/pkg/e"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}


func (g *Gin)FastReturnCode(httpCode, errCode int)  {
	g.FastReturn(httpCode,errCode,make(map[string]string))
}
func (g *Gin)FastReturn(httpCode, errCode int,data interface{})  {
	g.C.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})
}
