package env

import (
	"os"
	"log"
	
)

// GetEnv get key environment variable if exist otherwise return defalutValue
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		if defaultValue != "notSet" {
			value = defaultValue
		} else {
			errormessage := "ERROR: Environment Value " + key + " not set!"
			log.Fatal(errormessage)
		}
	}
	return value
}