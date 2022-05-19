package config

import "os"

const (
	envKey   = "APP_ENV"
	EnvLocal = "local"
	verKey   = "VERSION"
	VerLocal = "1.0.0"
)

var env = GetEnv(envKey, EnvLocal)
var VERSION = GetEnv(verKey, VerLocal)

func Env() string {
	return env
}

func GetEnv(key, def string) string {
	env, ok := os.LookupEnv(key)
	if ok {
		return env
	}

	return def
}
