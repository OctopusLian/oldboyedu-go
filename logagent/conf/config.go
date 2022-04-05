/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-04-05 13:46:17
 * @LastEditors: neozhang
 * @LastEditTime: 2022-04-05 14:40:17
 */
package conf

type AppConf struct {
	KafkaConf   `ini:"kafka"`
	TaillogConf `ini:"taillog"`
	EtcdConf    `ini:"etcd"`
}

type KafkaConf struct {
	Address     string `ini:"address"`
	Topic       string `ini:"topic"`
	ChanMaxSize int    `ini:"chanmaxsize"`
}

type EtcdConf struct {
	Address string `ini:"address"`
	Timeout int    `ini:"timeout"`
	Key     string `ini:"key"`
}

type TaillogConf struct {
	FileName string `ini:"filename"`
}
