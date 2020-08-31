package main

import (
	"fmt"
	"os"
)

const fallbackport string = "8080"

// Returns TCP port
func getPort() string {
	value := getenv("PORT", fallbackport)
	return fmt.Sprintf(":%s", value)
}

// Reads env var and returns fallback value if it is not set
func getenv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
