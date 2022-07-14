package application

import (
	"go-practice/infrastructure"
	"go-practice/models"
	"go-practice/routes"

	"gorm.io/gorm"
)

type ApplicationInitialize struct {
	database infrastructure.DatabaseSetup
	routes   routes.RouteInitializer
}

func NewApplicationInitialize(
	database infrastructure.DatabaseSetup,
	routes routes.RouteInitializer,
) *ApplicationInitialize {
	return &ApplicationInitialize{
		database: database,
		routes:   routes,
	}
}

func migrate(database *gorm.DB) {
	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.Product{})
	database.AutoMigrate(&models.ProductImage{})
	database.AutoMigrate(&models.Category{})
	database.AutoMigrate(&models.Comment{})
}

func (ra ApplicationInitialize) RunApplication() {
	db, _ := ra.database.SetupDatabase().DB()
	defer db.Close()
	migrate(ra.database.SetupDatabase())
	ra.routes.RouteSetup().Run()
}
