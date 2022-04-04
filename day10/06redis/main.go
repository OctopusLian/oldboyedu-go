package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

// redis

var redisdb *redis.Client

func initRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err = redisdb.Ping().Result()
	return
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Printf("connect redis failed, err:%v\n", err)
		return
	}
	fmt.Println("连接redis成功！")
	// zset
	key := "rank"
	items := []*redis.Z{
		&redis.Z{Score: 90, Member: "PHP"},
		&redis.Z{Score: 96, Member: "Golang"},
		&redis.Z{Score: 97, Member: "Python"},
		&redis.Z{Score: 99, Member: "Java"},
	}
	// 把元素都追加到key
	redisdb.ZAdd(key, items...)
	// 给Golang + 10分
	newScore, err := redisdb.ZIncrBy(key, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)
}
