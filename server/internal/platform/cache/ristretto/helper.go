package ristretto

import "fmt"

func (c *RistrettoCache) wrap(operation string, err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("ristretto.%s: %w", operation, err)
}
