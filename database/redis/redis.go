package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type RedisCli struct {
	conn redis.Conn
}

var instanceRedisCli *RedisCli = nil

func Init() {
	Connect()
	fmt.Println("redis initialized")
}
func Connect() (conn *RedisCli) {
	fmt.Println("redis conn")
	fmt.Println(instanceRedisCli)
	if instanceRedisCli == nil {
		fmt.Println("redis re initialized")
		instanceRedisCli = new(RedisCli)
		var err error

		instanceRedisCli.conn, err = redis.Dial("tcp", ":6003")

		if err != nil {
			panic(err)
		}
	}

	return instanceRedisCli
}

func (redisCli *RedisCli) SetValue(key string, value string, expiration ...interface{}) error {
	_, err := redisCli.conn.Do("SET", key, value)

	if err == nil && expiration != nil {
		redisCli.conn.Do("EXPIRE", key, expiration[0])
	}

	return err
}

func (redisCli *RedisCli) GetValue(key string) (interface{}, error) {
	return redisCli.conn.Do("GET", key)
}