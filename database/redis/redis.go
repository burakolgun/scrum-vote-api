package redis

import "github.com/go-redis/redis"

var Client *redis.Client
func Init()  {
	println("initialized -> Redis Client")
	Client = redis.NewClient(&redis.Options{
		Addr: "localhost:6003",
	})
}