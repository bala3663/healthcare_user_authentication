package main

import (
	"Final-Project-gin_helthcare/database"
	"Final-Project-gin_helthcare/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	port := "8000"
	router := gin.New() //router initialization
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.Patient_info_routes(router)
	routes.UserRoutes(router)
	router.Run(":" + port)

}
