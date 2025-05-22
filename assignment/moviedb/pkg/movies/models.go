package movies

import "github.com/google/uuid"

type Movie struct {
	Name   string    `json:"name"`
	Id     uuid.UUID `json:"id"`
	Actors []Actor   `json:"actors"`
}

type Actor struct {
	Name string `json:"name"`
}

type DeleteMovieReq struct {
	Name string `json:"name"`
}

type CreateMovieReq struct {
	Name   string  `json:"name"`
	Actors []Actor `json:"actors"`
}

type UpdateMovieReq struct {
	Name   string  `json:"name"`
	Actors []Actor `json:"actors"`
}
