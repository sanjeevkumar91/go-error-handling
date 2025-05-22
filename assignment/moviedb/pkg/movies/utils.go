package movies

import "github.com/google/uuid"

func removeByID(movies []Movie, idToRemove uuid.UUID) []Movie {
	var updatedMovies []Movie

	for _, movie := range movies {
		if movie.Id != idToRemove {
			updatedMovies = append(updatedMovies, movie)
		}
	}

	return updatedMovies
}

func updateByID(movies []Movie, idToUpdate uuid.UUID, updatedMovieReq UpdateMovieReq) []Movie {
	var updatedMovies []Movie

	for _, movie := range movies {
		if movie.Id == idToUpdate {
			updatedMovie := Movie{
				Id:     idToUpdate,
				Name:   updatedMovieReq.Name,
				Actors: updatedMovieReq.Actors,
			}
			updatedMovies = append(updatedMovies, updatedMovie)
		} else {
			updatedMovies = append(updatedMovies, movie)
		}
	}

	return updatedMovies
}
