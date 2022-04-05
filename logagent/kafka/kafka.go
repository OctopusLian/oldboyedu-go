/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-04-04 23:34:26
 * @LastEditors: neozhang
 * @LastEditTime: 2022-04-05 13:45:06
 */
package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// 专门往kafka写日志的模块

var (
	client sarama.SyncProducer // 声明一个全局的连接kafka的生产者client

)

// Init 初始化client
func Init(addrs []string) (err error) {
	config := sarama.NewConfig()
	// tailf包使⽤
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出⼀个 partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 连接kafka
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	return
}

func SendToKafka(topic, data string) {
	// 构造⼀个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)
	// 发送到kafka
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
