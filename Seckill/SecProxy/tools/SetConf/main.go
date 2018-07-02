package main

import (
	"github.com/coreos/etcd/clientv3"
	"time"
	"fmt"
	"encoding/json"
	"context"
)

const(
	EtcdKey = "/cjk/secskill/product"
)

//秒杀商品信息
type SecInfoConf struct {
	ProductId int
	StartTime int
	EndTime   int
	Status    int
	Total     int
	left      int
}


func SetLogConfToEtcd()  {
	cli,err := clientv3.New(clientv3.Config{
		Endpoints:[]string{"127.0.0.1:2379"},
		DialTimeout:5*time.Second,
	})
	if err != nil{
		fmt.Println("connect failed,err:%v",err)
	}
	fmt.Println("connect succ")
	defer cli.Close()

	var SecInfoConfArr []SecInfoConf
	SecInfoConfArr = append(SecInfoConfArr,
		SecInfoConf{
			ProductId:1022,
			StartTime:1517495965,
			EndTime:1517582365,
			Status:0,
			Total:10000,
			left:10000,
		},
    )

	SecInfoConfArr = append(SecInfoConfArr,
		SecInfoConf{
			ProductId:1052,
			StartTime:1517495965,
			EndTime:1517582365,
			Status:0,
			Total:900,
			left:900,
		},
	)
	data,err := json.Marshal(SecInfoConfArr)
	if err != nil{
		fmt.Println("json failed,err:%s",err)
		return
	}

	ctx,cancel := context.WithTimeout(context.Background(),time.Second)
	_,err = cli.Put(ctx,EtcdKey,string(data))
	cancel()
	if err != nil{
		fmt.Println("put failed ")
		return
	}

	ctx,cancel = context.WithTimeout(context.Background(),time.Second)
	resp,err := cli.Get(ctx,EtcdKey)
	cancel()

	if err != nil{
		fmt.Println("get failed err",err)
		return
	}
	for _,ev := range resp.Kvs{
		fmt.Printf("%s:%s\n",ev.Key,ev.Value)
	}
}
func main() {
    SetLogConfToEtcd()
}
