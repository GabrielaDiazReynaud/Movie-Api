package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine){
	server.GET("/movies", getAllMovies)
	server.GET("/movies/:id", getMovie)
	server.POST("/movies", saveMovie)
	server.PUT("/movies/:id", updateMovie)
	server.DELETE("/movies/:id", deleteMovie)
}