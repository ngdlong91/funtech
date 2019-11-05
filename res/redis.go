// Package res
package res

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v7"
)

type CRedis interface {
	Get(key string) ([]byte, error)
	Set(key string, val interface{}) error
	IsRun() bool
}

type cRedis struct {
	client      *redis.Client
	isConnected bool
}

func (r *cRedis) IsRun() bool {
	return r.isConnected
}

func (r *cRedis) Get(key string) ([]byte, error) {
	val, err := r.client.Get(key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
	}

	return []byte(val), nil
}

func (r *cRedis) Set(key string, val interface{}) error {
	if err := r.client.Set(key, val, 0).Err(); err != nil {
		return err
	}
	return nil
}

func doConnect() (*redis.Client, error) {
	server := os.Getenv("REDIS_SERVER")
	if len(server) == 0 {
		server = "localhost:6379"
	}
	client := redis.NewClient(&redis.Options{
		Addr:     server,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to redis")
	return client, nil
}

func RedisInstance() CRedis {
	r := &cRedis{
		isConnected: false,
	}

	client, err := doConnect()
	if err != nil {
		return r
	}
	r.isConnected = true
	r.client = client
	return r
}
