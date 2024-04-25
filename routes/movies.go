package routes

import (
	"net/http"
	"strconv"

	"example.com/movieApi/models"
	"github.com/gin-gonic/gin"
)

func getAllMovies(context *gin.Context){
	movies, err := models.GetAllMovies();
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not get events"})
		return
	}
	context.JSON(http.StatusOK, movies)
}

func getMovie(context * gin.Context){
	movieId, err := strconv.ParseInt( context.Param("id"), 10, 64)
	if err !=nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse movie ID"})
		return
	}
	movie, err:= models.GetMovie(movieId)
	if err !=nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not get movie"})
		return
	}
	context.JSON(http.StatusOK, movie)
}

func saveMovie(context *gin.Context){
	var movie models.Movie

	err := context.ShouldBindJSON(&movie)
	if err !=nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not get movie info"})
		return
	}
	err = movie.SaveNewMovie()
	if err !=nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not save Movie"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Movie created"})
}

func updateMovie(context *gin.Context){
	movieId, err := strconv.ParseInt( context.Param("id"), 10, 64)
	if err !=nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse movie ID"})
		return
	}

	_, err = models.GetMovie(movieId)
	if err !=nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not find Movie"})
		return
	}
	var newMovieData models.Movie
	err = context.ShouldBindJSON(&newMovieData)
	if err !=nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not get movie info"})
		return
	}
	newMovieData.ID = movieId

	err= newMovieData.UpdateMovie()
	if err !=nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not update movie"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Movie updated"})

}

func deleteMovie(context *gin.Context){
	movieId, err := strconv.ParseInt(context.Param("id"), 10,64)

	if err !=nil {
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse movie ID"} )
		return
	}

	movie, err := models.GetMovie(movieId)
	if err !=nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not find Movie"})
		return
	}
	err = movie.DeleteMovie()
	if err !=nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not delete movie"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Movie deleted"})


}