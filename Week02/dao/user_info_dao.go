package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"time"
)


//user表结构体定义
type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
}


//数据库连接信息
const (
	USERNAME = "root"
	PASSWORD = ""
	NETWORK = "tcp"
	SERVER = "127.0.0.1"
	PORT = 3306
	DATABASE = "test"
)


var (
	MysqlDB    *sql.DB
	MysqlDBErr error
)

func init()  {
	MysqlDB, MysqlDBErr = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if MysqlDBErr != nil {
		fmt.Println(MysqlDBErr)
		panic("conn mysql error")
	}
	MysqlDB.SetConnMaxLifetime(100*time.Second)  //最大连接周期，超时的连接就close
	MysqlDB.SetMaxOpenConns(100)                //设置最大连接数
	if MysqlDBErr = MysqlDB.Ping(); nil != MysqlDBErr {
		panic("数据库链接失败: " + MysqlDBErr.Error())
	}

}

func SelectUserById(id int) (*User,error) {
	row := MysqlDB.QueryRow("select id,name from user where  id = ?", id)
	user := new(User)
	err := row.Scan(&user.Id, &user.Username)
	if err != nil{
		return nil, errors.Wrapf(err,"user query error  uid :%v",id)
	}
	return user,nil
}
