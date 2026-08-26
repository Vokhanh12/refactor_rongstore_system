package serializer

import "fmt"

func wrap(operation string, err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("json.%s: %w", operation, err)
}
