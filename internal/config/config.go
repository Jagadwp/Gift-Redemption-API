package config

import (
    "fmt"
    "os"
    "strconv"
    "strings"
    "net/url"

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
	URL      string
}

type JWTConfig struct {
	Secret      string
	ExpiryHours int
}

func (d DatabaseConfig) DSN() string {
	// If DATABASE_URL exists (Heroku)
	if d.URL != "" {
		if strings.Contains(d.URL, "sslmode=") {
			return d.URL
		}

		sep := "?"
		if strings.Contains(d.URL, "?") {
			sep = "&"
		}

		return d.URL + sep + "sslmode=require"
	}

	// Local development DSN
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		d.Host,
		d.Port,
		d.User,
		d.Password,
		d.Name,
	)
}

func (d DatabaseConfig) MigrationURL() string {
    if d.URL != "" {
        if strings.Contains(d.URL, "sslmode=") {
            return d.URL
        }
        sep := "?"
        if strings.Contains(d.URL, "?") {
            sep = "&"
        }
        return d.URL + sep + "sslmode=require"
    }

    u := &url.URL{
        Scheme: "postgres",
        User:   url.UserPassword(d.User, d.Password),
        Host:   d.Host + ":" + d.Port,
        Path:   d.Name,
        RawQuery: "sslmode=disable",
    }
    return u.String()
}

func Load() *Config {
	appEnv := getEnv("APP_ENV", "development")

	// Only load .env in development
	if appEnv == "development" {
		_ = godotenv.Load()
	}

	jwtExpiry, _ := strconv.Atoi(getEnv("JWT_EXPIRY_HOURS", "24"))

	port := getEnv("PORT", "")
	if port == "" {
		port = getEnv("APP_PORT", "8080")
	}

	return &Config{
		AppPort: port,
		AppEnv:  appEnv,
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", "gift_redemption"),
			URL:      getEnv("DATABASE_URL", ""),
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
