package main

import (
	"fmt"
	"net"
)

func handleConn(c net.Conn) {
	defer c.Close()

	fmt.Println("准备读取...")
	for {
		var buf = make([]byte, 10)
		n, err := c.Read(buf)
		if err != nil {
			fmt.Println(fmt.Sprintf("读取发生错误：%+v", err))
			return
		}
		fmt.Println(fmt.Sprintf("已读取 %d 字节，内容是：%s", n, string(buf)))
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
