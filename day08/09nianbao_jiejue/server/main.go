/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-04-04 18:35:49
 * @LastEditors: neozhang
 * @LastEditTime: 2022-04-04 18:51:25
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"

	proto "oldboyedu-go/day08/09nianbao_jiejue/protocol"
)

// socket_stick/server/main.go

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		recvStr, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode failed,err:", err)
			return
		}
		fmt.Println("收到client发来的数据：", recvStr)
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}
