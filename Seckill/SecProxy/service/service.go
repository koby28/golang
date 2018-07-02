package service

import (
	"github.com/astaxie/beego/logs"
	"fmt"
	"time"
)

var (
	secKillConf *SecKillConf
)

func InitService(serviceConf *SecKillConf)  {
	secKillConf = serviceConf
	logs.Debug("init service succ,config:%v",secKillConf)
}

func SecInfoList() (data []map[string]interface{},code int,err error) {
	secKillConf.RwSecProductlock.RLock()
	defer secKillConf.RwSecProductlock.Unlock()
	for _,v := range secKillConf.SecProductInfoMap{
		item,_,err := SecInfoById(v.ProductId)
		if(err != nil){
			logs.Error("get product_id[%d] failed,err:%v",v.ProductId,err)
			continue
		}
        data = append(data,item)

	}
	return
}

func SecInfo(productId int)(data []map[string]interface{},code int,err error)  {
	secKillConf.RwSecProductlock.Lock()
	defer secKillConf.RwSecProductlock.Unlock()
    logs.Debug("sec info config is [%v]",secKillConf.SecProductInfoMap[productId])


	return
}

func SecInfoById(productId int)(data map[string]interface{},code int,err error)  {
	secKillConf.RwSecProductlock.Lock()
	defer secKillConf.RwSecProductlock.Unlock()

	v,ok := secKillConf.SecProductInfoMap[productId]

	if !ok {
		code = 1002
		err = fmt.Errorf("not found product_id:%d",productId)
		return
	}
	start := false
	end := false
	status := "success"

	now := time.Now().Unix()
	logs.Info(now,v.StartTime)
	if now - v.StartTime < 0 {
		start = false
		end = false
		status = "sec kill is not start"
	}

	if now - v.StartTime > 0 {
		start =  true
	}

	if now - v.EndTime > 0 {
		start = false
		end = true
		status = "sec kill is already end"
	}
    if v.Status == ProductStatusForceSaleOut || v.Status == ProductStatusSaleOut{
    	start = false
    	end = true
    	status = "product is sale out"
	}

	data = make(map[string]interface{})
	data["product_id"] = productId
	data["start_time"] = v.StartTime
	data["end_time"] = v.EndTime
	data["status"] = v.Status
	return
}
