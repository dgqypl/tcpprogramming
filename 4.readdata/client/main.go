package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("begin dial...")
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	defer conn.Close()
	fmt.Println("dial ok")

	var n int
	if n, err = conn.Write([]byte("hello world! Gopher")); err != nil { // 19 字节
		fmt.Println(fmt.Sprintf("写入发生错误：%+v", err))
		return
	}

	fmt.Println(fmt.Sprintf("已写入 %d 字节", n))
}
