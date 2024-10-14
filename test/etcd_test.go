package test

//
//import (
//	"context"
//	"fmt"
//	clientv3 "go.etcd.io/etcd/client/v3"
//	"testing"
//	"time"
//)
//
//func TestEtcd(t *testing.T) {
//	cli, err := clientv3.New(clientv3.Config{
//		Endpoints:   []string{"127.0.0.1:2379"},
//		DialTimeout: 5 * time.Second,
//	})
//
//	if err != nil {
//		// handle error!
//		fmt.Printf("connect to etcd failed, err:%v\n", err)
//		return
//	}
//	fmt.Println("connect to etcd success")
//	defer cli.Close()
//
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//	// put
//	//_, err = cli.Put(ctx, "test", fmt.Sprintf("test-val: %v", time.Now().Format(time.RFC3339)))
//	//cancel()
//	//if err != nil {
//	//	fmt.Printf("put to etcd failed, err:%v\n", err)
//	//	return
//	//}
//
//	// get
//	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
//	resp, err := cli.Get(ctx, "sys.rpc")
//	cancel()
//	if err != nil {
//		fmt.Printf("get from etcd failed, err:%v\n", err)
//		return
//	}
//	fmt.Println("resp Kvs", resp.Kvs)
//	for _, ev := range resp.Kvs {
//		fmt.Printf("xxx %s:%s\n", ev.Key, ev.Value)
//	}
//}
