package miniredis_demo

import (
	"context"
	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"testing"
	"time"
)

func TestDoSomethingWithRedis(t *testing.T) {
	// mock 一个redis server
	server, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer server.Close()

	// 准备数据
	server.Set("betuli", "yuay.ac.cn")
	server.SAdd(KeyValidWebsite, "betuli")

	// 连接mock的redis server
	rdb := redis.NewClient(&redis.Options{
		Addr: server.Addr(),
	})

	// 调用我们的函数
	ok := DoSomethingWithRedis(rdb, "betuli")
	if !ok {
		t.Fatal()
	}

	// 可以手动检查redis中的值是否复合预期
	if got, err := rdb.Get(context.TODO(), "blog").Result(); err != nil {
		t.Fatal(err)
	} else if got != "https://yuay.ac.cn" {
		t.Fatalf("unexpected value: %s", got)
	}

	// 也可以使用工具进行检查
	server.CheckGet(t, "blog", "https://yuay.ac.cn")

	// 过期检查
	server.FastForward(5 * time.Second) // 快进5秒
	if server.Exists("blog") {
		t.Fatal("'blog' should not have existed anymore")
	}
}
