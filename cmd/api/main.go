package main

import (
	"fmt"
	"time"

	app_services "duonglt.net/internal/application/services"
	inf_services "duonglt.net/internal/infrastructure/services"
	"duonglt.net/pkg/sf"
)

func main() {
	app_services.NewAuthService()

	j := inf_services.NewJWTService[int64](
		[]byte("secret"), 24*time.Hour,
	)
	tk, _ := j.Create(sf.New())
	id, err := j.Parse(tk)
	if err != nil {
		fmt.Printf("failed to parse token: %+v\n", err)
		return
	}
	println(*id)
	// viper.SetConfigFile(".env")
	// if err := viper.ReadInConfig(); err != nil {
	// 	log.Printf("failed to read config: %+v\n", err)
	// 	os.Exit(1)
	// }
	// dns := viper.GetString("SENTRY_DSN")
	// if dns != "" {
	// 	if err := st.Init(st.ClientOptions{
	// 		Dsn:           viper.GetString("SENTRY_DSN"),
	// 		EnableTracing: true,
	// 		Environment:   viper.GetString("SENTRY_ENVIRONMENT"),
	// 		// Set TracesSampleRate to 1.0 to capture 100%
	// 		// of transactions for performance monitoring.
	// 		// We recommend adjusting this value in production,
	// 		TracesSampleRate: 1.0,
	// 	}); err != nil {
	// 		log.Printf("failed to initialize sentry: %v\n", err)
	// 	}
	// }

	// r, err := internal.InitializeRouter()
	// if err != nil {
	// 	log.Printf("failed to initialize router: %+v\n", err)
	// 	os.Exit(1)
	// }
	// if err := r.ServeHTTP(); err != nil {
	// 	log.Printf("failed to serve http: %+v\n", err)
	// 	os.Exit(1)
	// }
}
