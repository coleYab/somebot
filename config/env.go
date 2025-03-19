package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	DbName     string
	DbPassword string
	DbUsername string
	DbPort     string
	DbHost     string
	DbSchema   string
	Port       string
	BotToken   string
}

var Envs = Config{
	DbName:     getEnv("DB_DATABASE", "some-db"),
	DbPassword: getEnv("DB_PASSWORD", "some-pwd"),
	DbUsername: getEnv("DB_USERNAME", "some-thing"),
	DbPort:     getEnv("DB_PORT", "some-thang"),
	Port:       getEnv("DB_PORT", ":8080"),
	DbHost:     getEnv("DB_HOST", "localhost"),
	DbSchema:   getEnv("DB_SCHEMA", "public"),
	BotToken:   getEnv("BOT_TOKEN", "bot token is required"),
}

func getEnv(name, fallback string) string {
	if val, ok := os.LookupEnv(name); ok {
		return val
	}

	return fallback
}
