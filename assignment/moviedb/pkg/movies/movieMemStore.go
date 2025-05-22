package movies

import (
	"errors"

	"github.com/google/uuid"
)

var (
	NotFoundErr = errors.New("not found")
)

type MemStore struct {
	list []Movie
}

func NewMemStore() *MemStore {
	list := make([]Movie, 0)
	return &MemStore{
		list,
	}
}

func (m *MemStore) Add(movie CreateMovieReq) error {
	newMovie := Movie{
		Id:     uuid.New(),
		Name:   movie.Name,
		Actors: movie.Actors,
	}
	m.list = append(m.list, newMovie)
	return nil
}

func (m *MemStore) Delete(id uuid.UUID) error {
	m.list = removeByID(m.list, id)
	return nil
}

func (m *MemStore) Update(id uuid.UUID, movie UpdateMovieReq) error {
	m.list = updateByID(m.list, id, movie)
	return nil
}

func (m *MemStore) List() ([]Movie, error) {
	return m.list, nil
}
