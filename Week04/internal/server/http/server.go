package http

import (
	"Go-000/Week04/internal/pkg/logging"
	"Go-000/Week04/internal/pkg/setting"
	v1 "Go-000/Week04/internal/server/http/v1"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func RunServer(ctx context.Context, serverPort int, serverName string) error {
	router := InitRouter()

	addr := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	logging.Info("server:",addr," start")
	s := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	defer func() {
		if err := recover(); err != nil {
			log.Printf("%+v",err)
			log.Println(serverName,"into defer recover()")
			cancelServer(s,serverName)
		}
	}()
	go func() {
		<-ctx.Done()
		cancelServer(s,serverName)
	}()
	log.Println(serverPort,serverName," with run")
	err := s.ListenAndServe()
	log.Println(serverName,"with error",err)
	return err
}

func cancelServer(s *http.Server, serverName string) {
	log.Println(serverName,"stop")
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(timeoutCtx)
}


func InitRouter() *gin.Engine {

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"message":"pong",
		})
	})
	router.GET("/test", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"message":"test",
		})
	})
	apiv1 := router.Group("/api/v1")
	//apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.UpdateTags)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		apiv1.GET("/GenerateArticlePoster",v1.GenerateArticlePoster)

	}

	return router
}

