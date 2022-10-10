package main

import (
	"fmt"
	"log"
	"net"
	"time"
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

	for i := 0; i < 2; i++ {
		time.Sleep(time.Duration(5) * time.Second)

		var n int

		if n, err = conn.Write([]byte("hello")); err != nil {
			fmt.Println(fmt.Sprintf("写入发生错误：%+v", err))
			return
		}

		fmt.Println(fmt.Sprintf("%s 已写入 %d 字节", time.Now().Format("2006-01-02 15:04:05.000"), n))
	}

}
