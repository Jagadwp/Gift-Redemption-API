package config

import (
    "fmt"
    "os"
    "strconv"

    "github.com/joho/godotenv"
)

type Config struct {
    AppPort  string
    AppEnv   string
    Database DatabaseConfig
    JWT      JWTConfig
}

type DatabaseConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    Name     string
}

type JWTConfig struct {
    Secret      string
    ExpiryHours int
}

func (d DatabaseConfig) DSN() string {
    return fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
        d.Host, d.Port, d.User, d.Password, d.Name,
    )
}

func Load() *Config {
    _ = godotenv.Load()

    jwtExpiry, _ := strconv.Atoi(getEnv("JWT_EXPIRY_HOURS", "24"))

    return &Config{
        AppPort: getEnv("APP_PORT", "8080"),
        AppEnv:  getEnv("APP_ENV", "development"),
        Database: DatabaseConfig{
            Host:     getEnv("DB_HOST", "localhost"),
            Port:     getEnv("DB_PORT", "5432"),
            User:     getEnv("DB_USER", "postgres"),
            Password: getEnv("DB_PASSWORD", ""),
            Name:     getEnv("DB_NAME", "gift_redemption"),
        },
        JWT: JWTConfig{
            Secret:      getEnv("JWT_SECRET", ""),
            ExpiryHours: jwtExpiry,
        },
    }
}

func getEnv(key, fallback string) string {
    if val := os.Getenv(key); val != "" {
        return val
    }
    return fallback
}