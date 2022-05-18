package common

import (
	"fmt"
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func setupDatabase(){
	db_host := os.Getenv("POSTGRES_HOST")
	db_port := os.Getenv("POSTGRES_PORT")
	db_user := os.Getenv("POSTGRES_USER")
	db_name := os.Getenv("POSTGRES_DB")
	db_password := os.Getenv("POSTGRES_PASSWORD")

	url := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v", db_host, db_port, db_user, db_name, db_password)
	db, _ := gorm.Open("postgres", url)
	
  	defer db.Close()
}