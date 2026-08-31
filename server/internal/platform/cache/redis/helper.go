package rediscache

import "fmt"

func (c *RedisCache) wrap(operation string, err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("redis.%s: %w", operation, err)
}
