package usecase

import (
	"context"

	"github.com/syafdia/sb/chal002/internal/domain/entity"
	"github.com/syafdia/sb/chal002/internal/domain/repo"
)

type FindMultipleMoviesUseCase interface {
	Execute(ctx context.Context, searchWord string, pagination int) ([]entity.MovieSummary, error)
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

func (f *findMultipleMoviesUseCase) Execute(ctx context.Context, searchWord string, pagination int) ([]entity.MovieSummary, error) {
	f.logRepo.Insert(ctx, map[string]interface{}{
		"action": "MultipleMovies",
		"payload": map[string]interface{}{
			"search_word": searchWord,
			"pagination":  pagination,
		},
	})

	return f.movieRepo.FindMultipleSummaries(ctx, searchWord, pagination)
}
