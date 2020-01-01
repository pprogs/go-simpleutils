package simpleutils

import "os"

//Getenv func
func Getenv(key, def string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return def
}
