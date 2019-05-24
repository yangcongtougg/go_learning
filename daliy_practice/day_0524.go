package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func createRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.107.131:6379",
		Password: "",
		DB:       0,
		PoolSize: 100, //连接池链接数
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return client
}

func setOperation(client *redis.Client) {
	client.SAdd("blacklist", "houshanjie")
	client.SAdd("blacklist", "wuyuchao")
	client.SAdd("blacklist", "zhangcong")
	client.SAdd("whitelist", "zhangcong")

	// 判断元素是否在集合中
	isMember, err := client.SIsMember("blacklist", "houshanjie").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("houshanjie is in the blacklist %v \n", isMember)

	// 求交集, 即既在黑名单中, 又在白名单中的元素
	names, err := client.SInter("blacklist", "whitelist").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("交集是%v \n", names)

	// 获取指定集合的所有元素
	all, err := client.SMembers("blacklist").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("黑名单的所有元素是%v \n", all)
}

func main() {
	client := createRedisClient()
	defer client.Close()

	setOperation(client)
}
