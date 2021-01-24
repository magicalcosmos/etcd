package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err: %v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()

	// cancel
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	values := `[{"path":"./nginx.log","topic":"web_log"},{"path":"./redis.log","topic":"redis_log"}]`
	_, err = cli.Put(ctx, "/logagent/192.168.1.6/collect_config", values)
	if err != nil {
		fmt.Printf("put to etcd failed, err: %v\n", err)
		return
	}
	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "sd")
	cancel()
	if err != nil {
		fmt.Printf("Get to etcd failed, err: %v\n", err)
		return
	}
	fmt.Println(resp)
	// ch := cli.Watch(context.Background(), "ll")
	// for wresp := range ch {
	// 	for _, evt := range wresp.Events {
	// 		fmt.Printf("Type: %v key: %v value %v\n", evt.Type, string(evt.Kv.Key), string(evt.Kv.Value))
	// 	}
	// }
}
