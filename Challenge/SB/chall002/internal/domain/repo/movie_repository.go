package repo

import (
	"context"

	"github.com/syafdia/sb/chal002/internal/domain/entity"
)

type MovieRepository interface {
	FindOne(ctx context.Context, imdbID string) (entity.Movie, error)
	FindMultiple(ctx context.Context, searchWord string, pagination int) ([]entity.Movie, error)
}
