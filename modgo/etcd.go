package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:32785"},
		// Endpoints: []string{"localhost:2379", "localhost:22379", "localhost:32379"}
		DialTimeout: 5 * time.Second,
	})

	cli.Do(clientv3.OpPut())
	aa, err := cli.Status(context.TODO(), "localhost:32785")
	fmt.Printf("%+v,%+v", aa, err)
}
