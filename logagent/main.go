/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-04-05 13:47:16
 * @LastEditors: neozhang
 * @LastEditTime: 2022-04-05 13:47:17
 */
package main

import (
	"fmt"
	"time"

	"code.oldboyedu.com/logagent/conf"
	"code.oldboyedu.com/logagent/kafka"
	"code.oldboyedu.com/logagent/taillog"
	"gopkg.in/ini.v1"
)

// logAgent入口程序

var (
	cfg = new(conf.AppConf)
)

func run() {
	// 1. 读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			// 2. 发送到kafka
			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// 0. 加载配置文件
	//cfg, err := ini.Load("./conf/config.ini")
	//fmt.Println(cfg.Section("kafka").Key("address"))
	//fmt.Println(cfg.Section("kafka").Key("topic"))
	//fmt.Println(cfg.Section("taillog").Key("path"))
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("load ini failed, err:%v\n", err)
		return
	}
	// 1. 初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Printf("init Kafka failed,err:%v\n", err)
		return
	}
	fmt.Println("init kafka success.")
	// 2. 打开日志文件准备收集日志
	err = taillog.Init(cfg.TaillogConf.FileName)
	if err != nil {
		fmt.Printf("Init taillog failed,err:%v\n", err)
		return
	}
	fmt.Println("init taillog success.")
	run()
}
