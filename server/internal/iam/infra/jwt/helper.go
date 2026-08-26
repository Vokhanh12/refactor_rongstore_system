package jwt

import "fmt"

func wrap(operation string, err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf(
		"jwt.%s: %w",
		operation,
		err,
	)
}
