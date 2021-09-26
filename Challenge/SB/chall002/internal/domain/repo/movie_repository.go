package repo

import (
	"context"

	"github.com/syafdia/sb/chal002/internal/domain/entity"
)

type MovieRepository interface {
	FindOneDetail(ctx context.Context, imdbID string) (entity.MovieDetail, error)
	FindMultipleSummaries(ctx context.Context, searchWord string, pagination int) ([]entity.MovieSummary, error)
}
