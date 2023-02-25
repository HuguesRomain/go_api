package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectToDB() (*gorm.DB, error) {
	host := "letsbookit.cbqzrg4ju3kj.eu-west-3.rds.amazonaws.com"
	port := "5432"
	user := "hugues"
	dbname := "letsbookit"
	password := "letsbookit"
	sslmode := "disable"

	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host, port, user, dbname, password, sslmode)

	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
