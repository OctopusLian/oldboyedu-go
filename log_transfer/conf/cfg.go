package conf

// LogTransferCfg 全局配置
type LogTransferCfg struct {
	KafkaCfg `ini:"kafka"`
	ESCfg `ini:"es"`
}

// KafkaCfg ...
type KafkaCfg struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

// ESCfg ...
type ESCfg struct {
	Address string `ini:"address"`
	ChanSize int `ini:"chan_size"`
	Nums int `ini:"nums"`
}
