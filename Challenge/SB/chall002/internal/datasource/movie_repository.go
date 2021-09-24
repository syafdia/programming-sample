package datasource

import (
	"context"
	"net/http"

	"github.com/syafdia/sb/chal002/internal/domain/entity"
	"github.com/syafdia/sb/chal002/internal/domain/repo"
)

type movieRepo struct {
	httpClient *http.Client
}

func NewMovieRepo(httpClient *http.Client) repo.MovieRepository {
	return &movieRepo{httpClient: httpClient}
}

func (m *movieRepo) FindOne(ctx context.Context, imdbID string) (entity.Movie, error) {
	return entity.Movie{}, entity.ErrNotImplemented
}
func (m *movieRepo) FindMultiple(ctx context.Context, searchWord string, pagination int) ([]entity.Movie, error) {
	return []entity.Movie{}, entity.ErrNotImplemented
}
