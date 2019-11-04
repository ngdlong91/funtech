// Package res
package res

import "testing"

func Test_redis_new(t *testing.T) {
	r := RedisInstance()
	t.Run("server not available", func(t *testing.T) {

	})

	t.Run("connect success", func(t *testing.T) {

	})

}

func Test_redis_get(t *testing.T) {
	t.Run("Key not available", func(t *testing.T) {

	})

	t.Run("Get key success", func(t *testing.T) {

	})
}

func Test_redis_set(t *testing.T) {
	t.Run("cannot set", func(t *testing.T) {

	})

	t.Run("set success", func(t *testing.T) {

	})
}
