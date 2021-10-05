package helpers

import (
	"os"
	"strconv"
)

const (
	TypeString = 1
	TypeInt    = 2
	TypeBool   = 3
	TypeDate   = 4
	TypeFloat  = 5
)

// EnvToInt convert env variable to int
func EnvToInt(value *int, key string, defaultValue int) {
	*value = defaultValue

	envValue, exists := os.LookupEnv(key)
	if !exists {
		return
	}

	if res, err := strconv.Atoi(envValue); err == nil {
		*value = res
	}
}

// EnvToBool convert env variable to bool
func EnvToBool(value *bool, key string, defaultValue bool) {
	*value = defaultValue

	envValue, exists := os.LookupEnv(key)
	if !exists {
		return
	}

	if res, err := strconv.ParseBool(envValue); err == nil {
		*value = res
	}
}

// EnvToString convert env variable to string
func EnvToString(value *string, key string, defaultValue string) {
	*value = getEnv(key, defaultValue)
}

// Simple helper function to read an environment or return a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
