package config

import (
	"os"
	"strconv"
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

type CassandraConfig struct {
	USER         string
	PASSWORD     string
	KEYSPACE     string
	CLUSTERS     []string
	APP_LABEL    string
	URL_START_ID int
	URL_END_ID   int
}

func GetAppConfig() *AppConfig {
	return &AppConfig{
		LISTEN_PORT_CREATE:  GetEnvOrPanic("LISTEN_PORT_CREATE"),
		LISTEN_PORT_FORWARD: GetEnvOrPanic("LISTEN_PORT_FORWARD"),
		FORWARD_DOMAIN:      GetEnvOrPanic("FORWARD_DOMAIN"),
		CREATE_DOMAIN:       GetEnvOrPanic("CREATE_DOMAIN"),
	}
}

func GetForwardDomain() string {
	return GetEnvOrPanic("FORWARD_DOMAIN")
}

func GetCreateDomain() string {
	return GetEnvOrPanic("CREATE_DOMAIN")
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

func GetCassandraConfig() *CassandraConfig {
	return &CassandraConfig{
		USER:         GetEnvOrPanic("CASSANDRA_USER"),
		PASSWORD:     GetEnvOrPanic("CASSANDRA_PASSWORD"),
		CLUSTERS:     strings.Split(GetEnvOrPanic("CASSANDRA_CLUSTERS"), ","),
		KEYSPACE:     GetEnvOrPanic("CASSANDRA_KEYSPACE"),
		APP_LABEL:    GetEnvOrPanic("CASSANDRA_APP_LABEL"),
		URL_START_ID: Str2IntOrPanic(GetEnvOrPanic("CASSANDRA_URL_START_ID")),
		URL_END_ID:   Str2IntOrPanic(GetEnvOrPanic("CASSANDRA_URL_END_ID")),
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

func Str2IntOrPanic(value string) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return i
}
