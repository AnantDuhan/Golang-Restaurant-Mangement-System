package routes

import (
	"github.com/gin-gonic/gin"
	controller "golang-restaurant-management-system/controllers"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controller.getFoods())
	incomingRoutes.GET("/foods':food_id", controller.getFood())
	incomingRoutes.POST("/foods", controller.CreateFood())
	incomingRoutes.PATCH("/foods/:food_id", controller.UpdateFood())

}