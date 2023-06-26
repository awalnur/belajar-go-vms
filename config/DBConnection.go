package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ORMConnection() (*gorm.DB, error) {
	dbConfig, err := DatabaseConfig()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	connString := "user=" + dbConfig.DatabaseUser +
		" password=" + dbConfig.DatabasePassword +
		" host=" + dbConfig.DatabaseHost +
		" port=" + dbConfig.DatabasePort +
		" dbname=" + dbConfig.DatabaseName +
		" sslmode=" + dbConfig.DatabaseSSLMode +
		" TimeZone=" + dbConfig.DatabaseTimeZone

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connString,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return db, nil

}
