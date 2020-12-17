package main

import (
	"Go-000/Week04/internal/data"
	"Go-000/Week04/internal/pkg/logging"
	"Go-000/Week04/internal/pkg/setting"
	simpleHttp "Go-000/Week04/internal/server/http"
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)


func main()  {

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		log.Println("global defer")
		cancel()
	}()

	setting.Setup()
	models.Setup()
	logging.Setup()

	g, _ := errgroup.WithContext(ctx)
	// server
	g.Go(func() error {
		return simpleHttp.RunServer(ctx,8001,"正式server")
	})
	//// debug server
	//g.Go(func() error {
	//	return runServer(ctx,8111,"debug server")
	//})
	c := make(chan os.Signal)
	signal.Notify(c,os.Interrupt)
	g.Go(func() error {
		return listenStopSignal(ctx,c,cancel)
	})
	log.Printf("Actual pid is %d", syscall.Getpid())
	if err := g.Wait(); err != nil {
		log.Println("g.wait finish")
		log.Printf("%+v",err)
	}
	log.Println("sever finish")
	time.Sleep(2 * time.Second)
}

func listenStopSignal(ctx context.Context, sig chan os.Signal, cancel context.CancelFunc) error {
	select {
	case sign := <- sig:
		log.Println("接受信号",sign)
		err := fmt.Errorf("handle signal: %d", sign)
		cancel()
		return err
	case <-ctx.Done():
		log.Println("listenStopSignal ctx.Done")
		return nil
	}

}






