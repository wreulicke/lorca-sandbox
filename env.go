package main

import (
	"os"
)

type fallback func() (string, error)

func envOrDefault(key string, f fallback) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return f()
	}
	return v, nil
}
