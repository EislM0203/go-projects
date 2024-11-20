package main

import (
	"github.com/gin-gonic/gin"
	"traunseenet.com/rest-api/db"
	"traunseenet.com/rest-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	
	routes.RegisterRoutes(server)
	server.Run(":8080")
	

}
