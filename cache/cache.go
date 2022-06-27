package cache

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func GetRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb

}

func SetAndGet() {
	rdb := GetRedisClient()

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist

}

func HSetAndHGet() {
	rdb := GetRedisClient()
	err := rdb.HSet(ctx, "hkey", "key1", "v1").Err()
	if err != nil {
		panic(err)
	}
	val, err := rdb.HGet(ctx, "hkey", "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val)
	}

}
func HSet(ctx context.Context, key string, values ...interface{}) error {
	rdb := GetRedisClient()
	return rdb.HSet(ctx, key, values...).Err()
}

func HGetResult(ctx context.Context, key, field string) (string, error) {
	rdb := GetRedisClient()
	return rdb.HGet(ctx, key, field).Result()
}
