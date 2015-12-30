package envious

import (
	"fmt"
	"os"
)

func ENV(variable string) string {
	result := os.Getenv(variable)
	if result == "" {
		panic(fmt.Sprintf("%s ENV variable is not set. Please set it by adding `%s=value` to the `.env` file.", variable, variable))
	}

	return result
}
