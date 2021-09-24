package usecase

import (
	"context"

	"github.com/syafdia/sb/chal002/internal/domain/entity"
	"github.com/syafdia/sb/chal002/internal/domain/repo"
)

type FindMultipleMoviesUseCase interface {
	Execute(ctx context.Context, searchWord string, pagination int) ([]entity.Movie, error)
}

type findMultipleMoviesUseCase struct {
	movieRepo repo.MovieRepository
	logRepo   repo.LogRepository
}

func NewFindMultipleMoviesUseCase(
	movieRepo repo.MovieRepository,
	logRepo repo.LogRepository,
) FindMultipleMoviesUseCase {
	return &findMultipleMoviesUseCase{movieRepo: movieRepo, logRepo: logRepo}
}

func (f *findMultipleMoviesUseCase) Execute(ctx context.Context, searchWord string, pagination int) ([]entity.Movie, error) {
	return []entity.Movie{}, entity.ErrNotImplemented
}
