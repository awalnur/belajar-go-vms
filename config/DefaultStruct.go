package config

import (
	"github.com/joho/godotenv"
	"os"
	"strings"
)

type Database struct {
	DatabaseUser     string
	DatabasePassword string
	DatabaseHost     string
	DatabasePort     string
	DatabaseName     string
	DatabaseSSLMode  string
	DatabaseTimeZone string
}

type BasePrivileges struct {
	Name     string
	NameCode string
}

func BasePrivilegesList() ([]BasePrivileges, error) {
	err := godotenv.Load()
	if err != nil {
		return []BasePrivileges{}, err
	}
	var valueMenu []BasePrivileges
	var baseMenu = strings.Split(os.Getenv("BASE_MENU_NAME"), ",")
	for _, value := range baseMenu {
		name := string(value)
		nameCode := strings.Replace(strings.ToLower(name), " ", "_", 1)
		valueMenu = append(valueMenu, BasePrivileges{
			Name:     name,
			NameCode: nameCode,
		})
	}
	return valueMenu, nil
}

func DatabaseConfig() (Database, error) {
	err := godotenv.Load()
	if err != nil {
		return Database{
			DatabaseUser:     "",
			DatabasePassword: "",
			DatabaseHost:     "",
			DatabasePort:     "",
			DatabaseName:     "",
			DatabaseSSLMode:  "",
			DatabaseTimeZone: "",
		}, err
	}
	return Database{
		DatabaseUser:     os.Getenv("DATABASE_USER"),
		DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
		DatabaseHost:     os.Getenv("DATABASE_HOST"),
		DatabasePort:     os.Getenv("DATABASE_PORT"),
		DatabaseSSLMode:  os.Getenv("DATABASE_NAME"),
		DatabaseTimeZone: os.Getenv("DATABASE_SSL_MODE"),
	}, nil
}
