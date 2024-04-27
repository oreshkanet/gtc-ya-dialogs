package config

import (
	"encoding/base64"
	"fmt"
	"os"
)

type Config struct {
	Domain        string
	OAuthClientID string
}

func LoadFromEnv() *Config {
	return &Config{
		Domain:        requireString("DOMAIN"),
		OAuthClientID: requireString("YANDEX_OAUTH_CLIENT_ID"),
	}
}

func requireBytes(name string) []byte {
	str := requireString(name)
	res, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		panic(fmt.Errorf("failed to decode var %v with base64: %w", err))
	}
	return res
}
func requireString(name string) string {
	res, ok := os.LookupEnv(name)
	if !ok {
		panic(fmt.Sprintf("required env var %s not found", name))
	}
	return res
}
