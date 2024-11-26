package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const (
	Young   = "Young_Obj"
	Old     = "Old_Obj"
	Station = 2
)

type Redis struct {
	db *redis.Client
	c  chan string
}

func Init() *Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	if rdb.Ping(context.Background()).Err() != nil {
		return &Redis{}
	}
	c := make(chan string)
	go func() {
		for v := range c {
			resp, _ := http.Get(v)
			os.MkdirAll(filepath.Dir("cache/"+v), os.ModePerm)
			fd, _ := os.Create("cache/" + v)
			io.Copy(fd, resp.Body)
			fd.Close()
			resp.Body.Close()
			rdb.HIncrBy(context.Background(), Old, v, 1)
		}
	}()
	return &Redis{
		db: rdb,
		c:  c,
	}
}
func (c *Redis) Nil() bool {
	return c.db == nil
}
func (c *Redis) Exists(key string) bool {
	if c.Nil() {
		return false
	}
	res, _ := c.db.HExists(context.Background(), Old, key).Result()
	return res
}
func (c *Redis) Incr(fd string) {
	if c.Nil() {
		return
	}
	//有序列表 增加
	val, _ := c.db.HIncrBy(context.Background(), Young, fd, 1).Result()
	// upgrade 策略
	if val > Station {
		c.Upgrade(fd)
	}
}
func (c *Redis) Upgrade(key string) {
	if c.Nil() {
		return
	}
	c.c <- key
}
