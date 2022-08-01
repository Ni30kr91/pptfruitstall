package db

import (
	"fmt"

	"github.com/go-redis/redis"
)

var Client *redis.Client

func Radis() {

	fmt.Println("Go Redis Tutorial")

	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := Client.Ping().Result()
	fmt.Println(pong, err)
	//fmt.Println("connected")
}
