package config

import (
	"github.com/eduspeak/eduspeak-backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm" 
	"os"
	"log"
	"github.com/joho/godotenv" 
) 
   
func goDotEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}
var dbPass string = goDotEnv("DBPASS")
var dbUser string = goDotEnv("DBUSER")
var dbName string = goDotEnv("DBNAME")

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

	Database.AutoMigrate(
		&models.Membership{},
		&models.Course{},
		&models.Article{},
		&models.Video{},
		&models.Quiz{},
		&models.Grade{},
		&models.Enroll{},
	)

	return nil
}