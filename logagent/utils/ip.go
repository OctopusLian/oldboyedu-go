/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-04-05 14:11:01
 * @LastEditors: neozhang
 * @LastEditTime: 2022-04-05 14:11:02
 */
package utils

import (
	"net"
	"strings"
)

// 获取本地对外IP
func GetOutboundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	//fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.IP.String(), ":")[0]
	return
}
