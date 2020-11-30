package main

import (
	"Go-000/Week02/service"
	"encoding/json"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"strconv"
)
func main()  {
	defer GlobalRecover()
	http.HandleFunc("/queryUser", queryUser)
	log.Fatalln(http.ListenAndServe(":8082", nil))

	////user, err := service.GetUser(-1)
	//user, err := service.GetUser(2)
	//if err != nil {
	//	log.Printf("query user %+v",err)
	//	return
	//}
	//log.Printf("user:%v",user)

}
func queryUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	query := r.URL.Query()
	id, err := strconv.Atoi(query.Get("uid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user,err := service.GetUser(id)
	if  err != nil {
		log.Printf("GetUser %+v",err)
		w.WriteHeader(http.StatusBadRequest)
		is := errors.Is(err, service.USER_DATA_ERROR)
		if is {
			w.Write([]byte(err.Error()))
		}
		return
	}
	printJson, err := json.Marshal(user)
	if err != nil {
		log.Printf("json format error %+v",err)
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(printJson)
}



func GlobalRecover() {
	if err := recover(); err != nil {
		log.Printf("GlobalRecover error %+v",err)
	}
}
