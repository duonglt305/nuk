package main

import (
	"log"
	"os"

	st "github.com/getsentry/sentry-go"
	"github.com/spf13/viper"

	"duonglt.net/internal/nuk"
)

func main() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("config file not found; skipping: %+v\n", err)
		} else {
			log.Printf("failed to read config: %+v\n", err)
			os.Exit(1)
		}
	}
	if err := st.Init(st.ClientOptions{
		Dsn:           viper.GetString("SENTRY_DSN"),
		EnableTracing: true,
		Environment:   viper.GetString("SENTRY_ENVIRONMENT"),
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	}); err != nil {
		log.Printf("failed to initialize sentry: %v\n", err)
	}

	r, err := nuk.InitializeRouter()
	if err != nil {
		log.Printf("failed to initialize router: %+v\n", err)
		os.Exit(1)
	}
	if err := r.ServeHTTP(); err != nil {
		log.Printf("failed to serve http: %+v\n", err)
		os.Exit(1)
	}
}
