package config

import (
	"errors"
	"os"
	"strconv"
	"time"
)

type Config struct {
	DataBaseURL  string
	HTTPPort     string
	SessionTTL   time.Duration
	CookieSecure bool
}

func Load() (Config, error) {
	dataBaseUrl := os.Getenv("DATABASE_URL")
	if dataBaseUrl == "" {
		return Config{}, errors.New("DATABASE_URL is requied")
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	sessionTTL := 24 * time.Hour

	if value := os.Getenv("SESSIONTTL"); value != "" {
		parsedSessionTTL, err := time.ParseDuration(value)

		if err != nil {
			return Config{}, err
		}

		sessionTTL = parsedSessionTTL
	}

	cookieSecure := false
	if value := os.Getenv("COOKIE_SECURE"); value != "" {
		parsedCookieSecure, err := strconv.ParseBool(value)

		if err != nil {
			return Config{}, err
		}

		cookieSecure = parsedCookieSecure
	}

	return Config{
		DataBaseURL:  dataBaseUrl,
		HTTPPort:     httpPort,
		SessionTTL:   sessionTTL,
		CookieSecure: cookieSecure,
	}, nil
}
