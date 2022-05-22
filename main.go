package main
import (
	"github.com/joho/godotenv"
	"go-practice/common"
	"go-practice/models"
)

func main(){
	godotenv.Load()

	db,_ := common.SetupDatabase().DB()
	defer db.Close()
	common.SetupDatabase().AutoMigrate(&models.User{})
}