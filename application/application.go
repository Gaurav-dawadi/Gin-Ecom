package application

import (
	"go-practice/infrastructure"
	"go-practice/models"
	"go-practice/routes"

	"gorm.io/gorm"
)

type ApplicationInitialize struct {
	database      infrastructure.DatabaseSetup
	routes        routes.RouteInitializer
	categoryRoute routes.CategoryRoutes
	commentRoute  routes.CommentRoutes
	userRoute     routes.UserRoutes
	productRoute  routes.ProductRoutes
}

func NewApplicationInitialize(
	database infrastructure.DatabaseSetup,
	categoryRoute routes.CategoryRoutes,
	commentRoute routes.CommentRoutes,
	productRoute routes.ProductRoutes,
	userRoute routes.UserRoutes,
	routes routes.RouteInitializer,
) *ApplicationInitialize {
	return &ApplicationInitialize{
		database:      database,
		categoryRoute: categoryRoute,
		commentRoute:  commentRoute,
		productRoute:  productRoute,
		userRoute:     userRoute,
		routes:        routes,
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
