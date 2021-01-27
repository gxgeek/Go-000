package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:10000")
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept error: %v\n", err)
			continue
		}
		// 开始goroutine监听连接
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	// 读写缓冲区
	rd := bufio.NewReader(conn)

	msg := make(chan string, 8)

	wr := bufio.NewWriter(conn)
	addr := conn.RemoteAddr()
	go doWriter(wr,addr ,msg)

	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			log.Printf("read error: %v\n", err)
			return
		}
		msg <- string(line)
	}
}

func doWriter(wr *bufio.Writer, addr net.Addr, ch <-chan string) {
	for s := range ch {
		wr.WriteString("receive "+addr.String()+":")
		wr.WriteString(s)
		wr.Flush()
	}

}
