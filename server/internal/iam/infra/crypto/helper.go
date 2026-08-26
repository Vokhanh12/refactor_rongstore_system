package crypto

import "fmt"

func wrap(operation string, err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf(
		"argon2id.%s: %w",
		operation,
		err,
	)
}
