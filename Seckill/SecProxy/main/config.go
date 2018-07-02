package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"fmt"
	"strings"
	"LearGoProject/Seckill/SecProxy/service"
)

var(
	secKillConf = &service.SecKillConf{
		SecProductInfoMap:make(map[int]*service.SecProductInfoConf,1024),
	}
)

func initConfig()(err error)  {
	redisAddr := beego.AppConfig.String("redis_addr")
	etcdAddr := beego.AppConfig.String("etcd_addr")

	logs.Debug("read config succ,redis addr:%v",redisAddr)
	logs.Debug("read config succ,etcd addr :%v",etcdAddr)

	secKillConf.RedisConf.RedisAddr = redisAddr
	secKillConf.EtcdConf.EtcdAddr = etcdAddr

	if len(redisAddr) == 0 || len(etcdAddr) == 0 {
		err = fmt.Errorf("init config failed,redis[%s] or etcd[%s] config is null",redisAddr,etcdAddr)
		return
	}

	redisMaxIdle,err := beego.AppConfig.Int("redis_max_idle")
	if err != nil{
		err = fmt.Errorf("init config failed,read redis_max_idle error:%v",err)
		return
	}

	redisMaxActive,err := beego.AppConfig.Int("redis_max_active")
	if err != nil {
		err = fmt.Errorf("init config failed,read redis_max_active:%v",err)
		return
	}

	redisIdleTImeOut,err := beego.AppConfig.Int("redis_idle_timeout")
	if err != nil{
		err = fmt.Errorf("init config failed,read redis_idle_timeout:%v")
		return
	}

	secKillConf.RedisConf.RedisMaxIdle=redisMaxIdle
	secKillConf.RedisConf.RedisMaxActive=redisMaxActive
	secKillConf.RedisConf.RedisIdleTimeout=redisIdleTImeOut

	etcdTimeout,err := beego.AppConfig.Int("etcd_timeout")
	if err != nil{
		err = fmt.Errorf("init config failed,read etcd_timeout error:%v",err)
		return
	}
	secKillConf.EtcdConf.Timeout = etcdTimeout
	secKillConf.EtcdConf.EtcdSecKeyPrefix = beego.AppConfig.String("etcd_sec_key_prefix")
	if len(secKillConf.EtcdConf.EtcdSecKeyPrefix) == 0{
		err = fmt.Errorf("init config failed, read etcd_sec_key error:%v",err)
		return
	}

	productKey := beego.AppConfig.String("etcd_product_key")
	if len(productKey) == 0{
		err = fmt.Errorf("init config failed, read etcd_product_key error:%v",err)
		return
	}

	//判断目录后缀是否包含"/"
	if strings.HasSuffix(secKillConf.EtcdConf.EtcdSecKeyPrefix,"/") ==  false{
		secKillConf.EtcdConf.EtcdSecKeyPrefix = secKillConf.EtcdConf.EtcdSecKeyPrefix+"/"
	}
	//拼接key
	secKillConf.EtcdConf.EtcdSecProductKey = fmt.Sprintf("%s%s",secKillConf.EtcdConf.EtcdSecKeyPrefix,productKey)

	secKillConf.LogPath = beego.AppConfig.String("log_path")
	secKillConf.Loglevel = beego.AppConfig.String("log_level")
	return
}

