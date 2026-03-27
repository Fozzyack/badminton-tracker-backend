package env

import "os"

func GetProduction() bool {
	env := os.Getenv("ENV")
	if env == "production" {
		return true
	}
	return false
}
