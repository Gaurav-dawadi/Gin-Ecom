package infrastructure

import (
	"fmt"
	"os"

	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseSetup struct{}

func NewDatabaseSetup() *DatabaseSetup {
	return &DatabaseSetup{}
}

func (ds DatabaseSetup) SetupDatabase() *gorm.DB {
	db_host := os.Getenv("POSTGRES_HOST")
	db_port := os.Getenv("POSTGRES_PORT")
	db_user := os.Getenv("POSTGRES_USER")
	db_name := os.Getenv("POSTGRES_DB")
	db_password := os.Getenv("POSTGRES_PASSWORD")
	db_sslmode := os.Getenv("SSL_MODE")
	db_timezone := os.Getenv("TIMEZONE")

	url := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=%v TimeZone=%v", db_host, db_port, db_user, db_name, db_password, db_sslmode, db_timezone)
	DB, _ := gorm.Open(postgres.Open(url), &gorm.Config{})

	return DB
}
