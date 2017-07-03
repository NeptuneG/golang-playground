package redis_player

import (
	"testing"
	"fmt"
)

func TestRedisPlayer(t *testing.T) {
	test, err := Get("test")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(test)
}