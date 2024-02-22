package main

import (
	"log"
	"os"

	st "github.com/getsentry/sentry-go"

	"duonglt.net/internal/nuk"
)

func main() {
	if err := st.Init(st.ClientOptions{
		Dsn:           "https://f56b7dbc9705eaeb3408b5a223437b85@o4504060703997952.ingest.sentry.io/4506790022414336",
		EnableTracing: true,
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
