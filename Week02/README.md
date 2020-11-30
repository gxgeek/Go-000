学习笔记

API--> service-> dao 
- dao层: 如果数据库连接出现 异常 这种情况如果是本案例中强依赖数据库
就直接panic 一个异常  。
当查询语句出现异常,这时候 因为dao是第一个发现这个错误的地方,所以我们应该wrap 这个异常,抛出去。
- ***service***: 代码运行dao 层的异常 这时候 我们可以直接将dao 层的异常直接返回,或则我们使用 wrapMessage 附属些其他信息，但不能
再次wrap 这个异常,不然堆栈会*2, 当运行获取到user信息 如果user信息不符合 需求 我们可以抛错,这时候由于
service 是第一个发现这个错误的地方所以这时候可以wrap这个异常。
- API层: 这个时候我们就可以定义一个全局global recover 来捕捉全局异常。保证程序不会挂掉。 

wrap: 可以保留堆栈信息,只用于当错误第一次出现的地方。

#### dao 
```go
func SelectUserById(id int) (*User,error) {
	row := MysqlDB.QueryRow("select id,name from user where  id = ?", id)
	user := new(User)
	err := row.Scan(&user.Id, &user.Username)
	if err != nil{
		return nil, errors.Wrapf(err,"user query error  uid :%v",id)
	}
	return user,nil
}

```

####service
```
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
```   

####API   

```go
    ***http serve r***     
-----------------------------
	defer GlobalRecover()
	http.HandleFunc("/queryUser", queryUser)
	log.Fatalln(http.ListenAndServe(":8082", nil))
    ——————————————————————————————————————
     main 启动
	////user, err := service.GetUser(-1)
	//user, err := service.GetUser(2)
	//if err != nil {
	//	log.Printf("query user %+v",err)
	//	return
	//}
	//log.Printf("user:%v",user)

```