// Package env provides utilities to work with environment variables.
package env

import "os"

// GetEnv feches an environment variable, and if doesn't exist, returns the default value.
func GetEnv(env, str string) string {
	if len(os.Getenv(env)) > 0 {
		return os.Getenv(env)
	}

	return str
}
