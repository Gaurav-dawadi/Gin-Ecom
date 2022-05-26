package application

import (
	"go-practice/infrastructure"
	"go-practice/models"
	"go-practice/routes"

	"gorm.io/gorm"
)

func migrate(database *gorm.DB) {
	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.Product{})
	database.AutoMigrate(&models.Category{})
	database.AutoMigrate(&models.Comment{})
}

func RunApplication() {
	db, _ := infrastructure.SetupDatabase().DB()
	defer db.Close()
	migrate(infrastructure.SetupDatabase())
	routes.RouteSetup().Run()
}
