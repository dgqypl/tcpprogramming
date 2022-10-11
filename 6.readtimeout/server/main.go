package main

import (
	"fmt"
	"net"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()

	fmt.Println("准备读取...")

	for {
		var buf = make([]byte, 10)
		c.SetDeadline(time.Now().Add(3 * time.Second))
		n, err := c.Read(buf)
		if err != nil {
			fmt.Println(fmt.Sprintf("读取发生错误：%+v", err))
			if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
				fmt.Println(fmt.Sprintf("%s ", time.Now().Format("2006-01-02 15:04:05.000")), "发生读超时")
				continue
			}
			return
		}
		fmt.Println(fmt.Sprintf("%s 已读取 %d 字节", time.Now().Format("2006-01-02 15:04:05.000"), n))
	}

}

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	fmt.Println("已启动 tcp server")

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}
		fmt.Println("接收到一个新的连接")
		// start a new goroutine to handle
		// the new connection.
		go handleConn(c)
	}
}
