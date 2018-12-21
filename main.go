package main

import (
	"flag"
	"os"
)

func main() {
	commit := flag.String("commit", "", "Git commit ID")
	branch := flag.String("branch", "", "Git branch")
	authority := flag.String("authority", getAuthority(), "The authority that generated the version")

	flag.Parse()
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}