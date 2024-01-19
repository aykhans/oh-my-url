package config

import (
	"os"
	"strings"
)

type DBName string

const (
	Postgres  DBName = "postgres"
	MongoDB   DBName = "mongodb"
	Cassandra DBName = "cassandra"
)

type AppConfig struct {
	LISTEN_PORT_CREATE  string
	LISTEN_PORT_FORWARD string
	FORWARD_DOMAIN      string
	CREATE_DOMAIN       string
}

type PostgresConfig struct {
	USER     string
	PASSWORD string
	HOST     string
	PORT     string
	DBNAME   string
}

func GetAppConfig() *AppConfig {
	return &AppConfig{
		LISTEN_PORT_CREATE:  GetEnvOrPanic("LISTEN_PORT_CREATE"),
		LISTEN_PORT_FORWARD: GetEnvOrPanic("LISTEN_PORT_FORWARD"),
		FORWARD_DOMAIN:      GetEnvOrPanic("FORWARD_DOMAIN"),
		CREATE_DOMAIN:       GetEnvOrPanic("CREATE_DOMAIN"),
	}
}

func GetPostgresConfig() *PostgresConfig {
	return &PostgresConfig{
		USER:     GetEnvOrPanic("POSTGRES_USER"),
		PASSWORD: GetEnvOrPanic("POSTGRES_PASSWORD"),
		HOST:     GetEnvOrPanic("POSTGRES_HOST"),
		PORT:     GetEnvOrDefault("POSTGRES_PORT", "5432"),
		DBNAME:   GetEnvOrPanic("POSTGRES_DB"),
	}
}

func GetDB() DBName {
	dbName := strings.ToLower(GetEnvOrPanic("DB"))
	switch dbName {
	case "postgres":
		return Postgres
	case "mongodb":
		return MongoDB
	case "cassandra":
		return Cassandra
	default:
		panic("Unknown database")
	}
}

func GetEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetEnvOrPanic(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic("Environment variable " + key + " is not set")
	}
	return value
}
