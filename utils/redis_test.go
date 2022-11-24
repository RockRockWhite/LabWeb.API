package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"testing"
)

func TestGetReidsClient(t *testing.T) {

	hits, err := GetReidsClient().Incr(context.Background(), "hits").Result()
	if err == redis.Nil {
		fmt.Println("nil")
	}

	fmt.Println(hits)
}
