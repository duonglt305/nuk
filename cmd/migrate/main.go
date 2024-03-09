package main

import (
	"fmt"
	v4 "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("failed to read config: %+v\n", err)
		os.Exit(1)
	}
	args := os.Args[1:]
	if len(args) < 1 {
		log.Println("no command provided")
		os.Exit(1)
	}
	command := args[0]
	switch command {
	case "up":
		up()
	case "down":
		down()
	default:
		fmt.Printf("unknown command: %s; available commands: up, down\n", command)
	}
}

func initial() *v4.Migrate {
	m, err := v4.New("file://db/migrations", viper.GetString("DATABASE_URL"))
	if err != nil {
		log.Printf("failed to initalize migration: %+v\n", err)
		os.Exit(1)
	}
	return m
}

func up() {
	m := initial()
	if err := m.Up(); err != nil {
		log.Printf("failed to apply up migration: %+v\n", err)
		os.Exit(1)
	}
}

func down() {
	m := initial()
	if err := m.Down(); err != nil {
		log.Printf("failed to apply down migration: %+v\n", err)
		os.Exit(1)
	}
}
