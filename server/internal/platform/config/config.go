package config

import (
	"log"
	"os"
	"strconv"

	"github.com/lpernett/godotenv"
)

type Config struct {

	// DB
	PostgresHost     string `envconfig:"POSTGRES_HOST"`
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PostgresMaxOpen  int
	PostgresMaxIdle  int

	// Redis
	RedisHost         string
	RedisPort         int
	RedisPassword     string
	RedisDB           int
	RedisTTL          int
	RedisPoolSize     int
	RedisMinIdleConns int

	// Keycloak
	KeycloakURL                     string
	KeycloakServerHealth            string
	KeycloakRealm                   string
	KeycloakClientID                string
	KeycloakSecret                  string
	KeycloakRedirect                string
	KeycloakScope                   string
	KeycloakGrantUmaTicketType      string
	KeycloakAudience                string
	KeycloakResponsePermissionsMode string

	// Title-gl
	TitleGlHost string
	TitleGlPort int

	// ORSM
	ORSMHost string
	ORSMPort string

	// RabitMQ
	RabbitMQHost     string
	RabbitMQPort     string
	RabbitMQUser     string
	RabbitMQPassword string

	// Grafana tempo
	TempoHost             string
	TempoPort             int
	TraceSampleRatio      float64
	DebugTraceSampleRatio float64

	// r2
	R2AccountID string
	R2AccessKey string
	R2SecretKey string
	R2Bucket    string
	R2CDNURL    string

	MaxRetries int
	Interval   int
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, reading env variables from system")
	}

	port, err := strconv.Atoi(getEnv("POSTGRES_PORT", "5432"))
	if err != nil {
		log.Fatalf("Invalid POSTGRES_PORT: %v", err)
	}

	postgresMaxOpen, err := strconv.Atoi(getEnv("POSTGRES_MAX_OPEN", "20"))
	if err != nil {
		log.Fatalf("Invalid REDIS_PORT: %v", err)
	}

	postgresMaxIdle, err := strconv.Atoi(getEnv("POSTGRES_MAX_IDLE", "10"))
	if err != nil {
		log.Fatalf("Invalid REDIS_PORT: %v", err)
	}

	redisPort, err := strconv.Atoi(getEnv("REDIS_PORT", "6379"))
	if err != nil {
		log.Fatalf("Invalid REDIS_PORT: %v", err)
	}

	redisDB, err := strconv.Atoi(getEnv("REDIS_DB", "0"))
	if err != nil {
		log.Fatalf("Invalid REDIS_DB: %v", err)
	}

	redisTTL, err := strconv.Atoi(getEnv("REDIS_TTL", "900"))
	if err != nil {
		log.Fatalf("Invalid REDIS_TTL: %v", err)
	}

	redisPoolSize, err := strconv.Atoi(getEnv("REDIS_POOL_SIZE", "20"))
	redisMinIdleConns, err := strconv.Atoi(getEnv("REDIS_MIN_IDLE_CONNS", "5"))

	tempoPort, err := strconv.Atoi(getEnv("TEMPO_PORT", "4317"))
	if err != nil {
		log.Fatalf("Invalid TEMPO_PORT: %v", err)
	}

	maxRetries, err := strconv.Atoi(getEnv("MAX_RETRIES", "3"))
	if err != nil {
		log.Fatalf("Invalid MAX_RETRIES: %v", err)
	}

	interval, err := strconv.Atoi(getEnv("INTERVAL", "1"))
	if err != nil {
		log.Fatalf("Invalid INTERVAL: %v", err)
	}

	return &Config{
		PostgresHost:     getEnv("POSTGRES_HOST", "localhost"),
		PostgresPort:     port,
		PostgresUser:     getEnv("POSTGRES_USER", "postgres"),
		PostgresPassword: getEnv("POSTGRES_PASSWORD", ""),
		PostgresDB:       getEnv("POSTGRES_DB", "postgres"),
		PostgresMaxOpen:  postgresMaxOpen,
		PostgresMaxIdle:  postgresMaxIdle,

		RedisHost:         getEnv("REDIS_HOST", "localhost"),
		RedisPort:         redisPort,
		RedisPassword:     getEnv("REDIS_PASSWORD", ""),
		RedisDB:           redisDB,
		RedisTTL:          redisTTL,
		RedisPoolSize:     redisPoolSize,
		RedisMinIdleConns: redisMinIdleConns,

		KeycloakURL:                     getEnv("KEYCLOAK_SERVER_URL", ""),
		KeycloakServerHealth:            getEnv("KEYCLOAK_SERVER_HEALTH", ""),
		KeycloakRealm:                   getEnv("KEYCLOAK_REALM", ""),
		KeycloakClientID:                getEnv("KEYCLOAK_CLIENT_ID", ""),
		KeycloakSecret:                  getEnv("KEYCLOAK_CLIENT_SECRET", ""),
		KeycloakRedirect:                getEnv("KEYCLOAK_REDIRECT_URI", ""),
		KeycloakScope:                   getEnv("KEYCLOAK_SCOPE", ""),
		KeycloakGrantUmaTicketType:      getEnv("KEYCLOAK_GRANT_UMA_TICKET_TYPE", ""),
		KeycloakAudience:                getEnv("KEYCLOAK_AUDIENCE", ""),
		KeycloakResponsePermissionsMode: getEnv("KEYCLOAK_RESPONSE_PERMISSIONS_MODE", ""),

		RabbitMQHost:     getEnv("RABBITMQ_HOST", ""),
		RabbitMQPort:     getEnv("RABBITMQ_PORT", ""),
		RabbitMQUser:     getEnv("RABBITMQ_USER", ""),
		RabbitMQPassword: getEnv("RABBITMQ_PASSWORD", ""),

		TempoHost:             getEnv("TEMPO_HOST", ""),
		TempoPort:             tempoPort,
		TraceSampleRatio:      mustParseFloat("TRACE_SAMPLE_RATIO", "0.05"),
		DebugTraceSampleRatio: mustParseFloat("DEBUG_TRACE_SAMPLE_RATIO", "1.0"),

		R2AccountID: getEnv("R2_ACCOUNT_ID", ""),
		R2AccessKey: getEnv("R2_ACCESS_KEY", ""),
		R2SecretKey: getEnv("R2_SECRET_KEY", ""),
		R2Bucket:    getEnv("R2_BUCKET", ""),
		R2CDNURL:    getEnv("R2_CDN_URL", ""),

		MaxRetries: maxRetries,
		Interval:   interval,
	}
}

func mustParseInt(key, def string) int {
	v, err := strconv.Atoi(getEnv(key, def))
	if err != nil {
		log.Fatalf("Invalid %s: %v", key, err)
	}
	return v
}

func mustParseFloat(key, def string) float64 {
	v, err := strconv.ParseFloat(getEnv(key, def), 64)
	if err != nil {
		log.Fatalf("Invalid %s: %v", key, err)
	}
	return v
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
