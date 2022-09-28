
package utils

import (
	"os"
)

func IsEnvExist(key string) bool {
	if _, ok := os.LookupEnv(key); ok {
			return true
	}
	
	return false
}