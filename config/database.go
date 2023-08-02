package config

import (
	"github.com/eduspeak/eduspeak-backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm" 
	"os"
	// "fmt"
) 

type Config struct {
	dbUser string
	dbPassword string
}
var dbPass string = os.Getenv("DBPASS")
var dbUser string = os.Getenv("DBUSER")
var dbName string = os.Getenv("DBNAME")

var Database *gorm.DB

var DATABASE_URI string = dbUser+":"+dbPass+"@tcp(localhost:3306)/"+dbName+"?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() error {
	var err error

	Database, err = gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt: true,
	})

	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&models.Membership{})

	return nil
}