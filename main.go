package main

import (
	"example.com/movieApi/db"
	"example.com/movieApi/routes"
	"github.com/gin-gonic/gin"
)
func main(){
	db.DBInit();
	server:= gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}