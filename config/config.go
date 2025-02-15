package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port                 string
	MongoUri             string
	DbName               string
	JwtSecret            string
	JwtExpirationMinutes int
	JwtKey               []byte
	MessagesPageSize     int64
	TinodeHttpHost       string
	TinodeWsHost         string
}

var AppConfig Config

func LoadConfig() {
	loadEnv()

	jwtExpirationMinutes, err := strconv.Atoi(getEnvWithFallback("JWT_EXPIRATION_MINUTES", "360"))
	if err != nil {
		panic(err)
	}

	messagesPageSize, err := strconv.ParseInt(getEnvWithFallback("MESSAGES_PAGE_SIZE", "50"), 10, 64)
	if err != nil {
		panic(err)
	}

	AppConfig = Config{
		Port:                 getEnvWithFallback("PORT", "8080"),
		MongoUri:             getEnvWithFallback("MONGO_URI", "mongodb://localhost:28017"),
		DbName:               getEnvWithFallback("DB_NAME", "tinode_chat"),
		JwtSecret:            getEnvWithFallback("JWT_SECRET", "tinode"),
		JwtExpirationMinutes: jwtExpirationMinutes,
		JwtKey:               []byte(getEnvWithFallback("JWT_SECRET", "tinode")),
		MessagesPageSize:     messagesPageSize,
		TinodeHttpHost:       getEnvWithFallback("TINODE_HTTP_HOST", "localhost:16060"),
		TinodeWsHost:         getEnvWithFallback("TINODE_WS_HOST", "localhost:16060"),
	}
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func getEnvWithFallback(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
