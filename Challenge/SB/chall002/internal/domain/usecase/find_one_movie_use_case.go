package usecase

import (
	"context"

	"github.com/syafdia/sb/chal002/internal/domain/entity"
	"github.com/syafdia/sb/chal002/internal/domain/repo"
)

type FindOneMovieUseCase interface {
	Execute(ctx context.Context, id string) (entity.MovieDetail, error)
}

type findOneMovieUseCase struct {
	movieRepo repo.MovieRepository
	logRepo   repo.LogRepository
}

func NewFindOneMovieUseCase(
	movieRepo repo.MovieRepository,
	logRepo repo.LogRepository,
) FindOneMovieUseCase {
	return &findOneMovieUseCase{movieRepo: movieRepo, logRepo: logRepo}
}

func (f *findOneMovieUseCase) Execute(ctx context.Context, id string) (entity.MovieDetail, error) {
	f.logRepo.Insert(ctx, map[string]interface{}{
		"action": "FindOneMovie",
		"payload": map[string]interface{}{
			"id": id,
		},
	})

	return f.movieRepo.FindOneDetail(ctx, id)
}
