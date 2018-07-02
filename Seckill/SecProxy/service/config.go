package service

import "sync"

type RedisConf struct {
	RedisAddr string
	RedisMaxIdle int
	RedisMaxActive int
	RedisIdleTimeout int
}

type EtcdConf struct {
	EtcdAddr string
	Timeout int
	EtcdSecKeyPrefix string
	EtcdSecProductKey string
}

type SecKillConf struct {
	RedisConf RedisConf
	EtcdConf EtcdConf
	LogPath string
	Loglevel string
	SecProductInfoMap map[int]*SecProductInfoConf
	RwSecProductlock sync.RWMutex
}

//秒杀商品信息
type SecProductInfoConf struct {
	ProductId int
	StartTime int64
	EndTime   int64
	Status    int
	Total     int
	Left      int
}

