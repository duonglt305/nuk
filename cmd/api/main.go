package main

import (
	"log"
	"os"

	"duonglt.net/internal"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("failed to read config: %+v\n", err)
		os.Exit(1)
	}
	r, err := internal.Initialize()
	if err != nil {
		log.Printf("%+v\n", err)
		os.Exit(1)
	}
	if err := r.ServeHTTP(); err != nil {
		log.Printf("failed to start application: %+v\n", err)
		os.Exit(1)
	}
}
