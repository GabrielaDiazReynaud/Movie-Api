package models

import (
	"example.com/movieApi/db"
)
type Movie struct {
	ID int64
	MovieName string `binding:"required"`
	Year int64	`binding:"required"`
	Duration string	`binding:"required"`
	Director string	`binding:"required"`
	Cast string `binding:"required"`
}

func GetAllMovies()([]Movie, error){
	query:= "SELECT * FROM movies"
	rows, err := db.DB.Query(query)
	if err !=nil{
		return nil, err
	}
	var movies []Movie

	for rows.Next(){
		var movie Movie
		err = rows.Scan(&movie.ID, &movie.MovieName, &movie.Year, &movie.Duration, &movie.Director, &movie.Cast)
		if err !=nil{
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, err
}

func GetMovie(movieId int64)(Movie, error){
	query:= "SELECT * FROM movies WHERE id= ?"
	row := db.DB.QueryRow(query, movieId)
	var movie Movie
	err :=row.Scan(&movie.ID, &movie.MovieName, &movie.Year, &movie.Duration, &movie.Director, &movie.Cast)
	if err !=nil{
		return Movie{}, err
	}
	return movie, nil
}

func (movie Movie) SaveNewMovie()error{
	query:= "INSERT INTO movies (movie_name, year, duration, director, cast) VALUES (?, ?, ?, ?, ?)"

	_, err:= db.DB.Exec(query, movie.MovieName, movie.Year, movie.Duration, movie.Director, movie.Cast)
	if err !=nil {
		return err
	}
	return nil
}

func (movie Movie) UpdateMovie()error{
	query:= `
	UPDATE movies
	SET movie_name = ?, year = ?, duration = ?, director = ?, cast= ?
	WHERE id = ? 
	`
	_, err:= db.DB.Exec(query, movie.MovieName, movie.Year, movie.Duration, movie.Director, movie.Cast, movie.ID)
	if err !=nil {
		return err
	}
	return nil
}

func (movie Movie) DeleteMovie() error{
	query:= "DELETE FROM movies WHERE id= ?"

	_, err := db.DB.Exec(query, movie.ID)
	if err !=nil {
		return err
	}
	return nil

}

