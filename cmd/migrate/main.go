package main

import (
	"fmt"
	"log"
	"os"

	"duonglt.net/pkg/db"
	v4 "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
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
	var dbDriver database.Driver
	var dbIns db.IRds
	var err error
	driver := viper.GetString("DB_DRIVER")
	dbUrl := viper.GetString("DB_URL")
	if driver == "" {
		log.Println("DB_DRIVER is not set")
		os.Exit(1)
	}
	path := fmt.Sprintf("file://db/migrations/%s", driver)
	dbIns, err = db.New(driver, dbUrl)
	if err != nil {
		log.Printf("failed to connect to database: %+v\n", err)
		os.Exit(1)
	}
	switch driver {
	case db.MYSQLDriver:
		dbDriver, err = mysql.WithInstance(dbIns.Get().DB, &mysql.Config{})
	case db.PGSQLDriver:
		dbDriver, err = postgres.WithInstance(dbIns.Get().DB, &postgres.Config{})
	}
	if err != nil {
		log.Printf("failed to initial database driver: %+v\n", err)
		os.Exit(1)
	}
	m, err := v4.NewWithDatabaseInstance(path, "postgres", dbDriver)
	if err != nil {
		log.Printf("failed to initial migration: %+v\n", err)
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
