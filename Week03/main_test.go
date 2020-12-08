package main

import (
	"context"
	"fmt"
	"time"
)






//
//func main() {
//	var once sync.Once
//	for i := 0; i < 5; i++ {
//		go func(i int) {
//			fun1 := func() {
//				fmt.Printf("i:=%d\n", i)
//			}
//			once.Do(fun1)
//		}(i)
//	}
//	// 为了防止主goroutine直接运行完了，啥都看不到
//	time.Sleep(50 * time.Millisecond)
//}

//func main() {
//	var once sync.Once
//	fun1 := func() {
//		fmt.Println("第一次打印")
//	}
//	once.Do(fun1)
//
//	fun2 := func() {
//		fmt.Println("第二次打印")
//	}
//
//	once.Do(fun2)
//}


func MainTest() {
	//ctx, cancel := context.WithCancel(context.Background())
	//go watch("【监控1】",ctx)
	//go watch(ctx,"【监控2】")
	//go watch(ctx,"【监控3】")
	//ctx =context.(ctx,"key","【22222】")
	//
	//time.Sleep(10 * time.Second)
	//fmt.Println("可以了，通知监控停止")
	//ctx =context.WithValue(ctx,"key","【333333】")
	//cancel()
	//ctx =context.WithValue(ctx,"key","【44444444】")
	////为了检测监控过是否停止，如果没有监控输出，就表示停止了
	//time.Sleep(5 * time.Second)
}
func watch( name string,ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name,"监控退出，停止了...",ctx.Value("key"))
			return
		default:
			fmt.Println(name,"goroutine监控中...",ctx.Value("key"))
			time.Sleep(2 * time.Second)
		}
	}
}


var Count int = 0
func main1()  {


		//g := new(errgroup.Group)
		//var urls = []string{
		//	"http://www.golang.org/",
		//	"http://www.google.com/",
		//	"http://www.somestupidname.com/",
		//}
		//for _, url := range urls {
		//	// Launch a goroutine to fetch the URL.
		//	url := url // https://golang.org/doc/faq#closures_and_goroutines
		//	g.Go(func() error {
		//		// Fetch the URL.
		//		resp, err := http.Get(url)
		//		if err == nil {
		//			resp.Body.Close()
		//		}
		//		return err
		//	})
		//}
		//// Wait for all HTTP fetches to complete.
		//if err := g.Wait(); err == nil {
		//
		//	fmt.Println("Successfully fetched all URLs.")
		//}else {
		//	fmt.Println(g.Wait().Error())
		//}

	//for i:= 0; i <100;i++ {
	//	CountAdd()
	//}
	//value := atomic.Value{}
	//value.
	//time.Sleep(100 * time.Millisecond)
	//fmt.Println(Count)
}

func CountAdd() {
	go func() {
		time.Sleep(2 * time.Millisecond)
		Count = Count+1
	}()
}
