package datasource

import (
	"context"

	"github.com/syafdia/sb/chal002/internal/domain/entity"
)

type logRepo struct {
}

func (l *logRepo) Insert(ctx context.Context, data entity.LogData) error {
	return entity.ErrNotImplemented
}
