/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-04-05 13:46:17
 * @LastEditors: neozhang
 * @LastEditTime: 2022-04-05 14:06:21
 */
package conf

type AppConf struct {
	KafkaConf   `ini:"kafka"`
	TaillogConf `ini:"taillog"`
}

type KafkaConf struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type EtcdConf struct {
	Address string `ini:"address"`
	Timeout int `ini:"timeout"`
}

type TaillogConf struct {
	FileName string `ini:"filename"`
}
