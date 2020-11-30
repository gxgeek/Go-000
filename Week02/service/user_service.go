package service

import (
	"Go-000/Week02/dao"
	"github.com/pkg/errors"
)

var USER_DATA_ERROR = errors.New("USER_DATA_ERROR ")

func GetUser(Id int) (*dao.User,error) {
	user, err := dao.SelectUserById(Id)
	if err != nil{
		return nil,errors.WithMessage(err,"not found user with service")
	}
	if user.Username == "" {
		return nil,errors.Wrap(USER_DATA_ERROR,"用户数据异常 请检查数据库")
	}
	return user,nil
}