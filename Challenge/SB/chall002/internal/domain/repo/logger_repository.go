package repo

import (
	"context"

	"github.com/syafdia/sb/chal002/internal/domain/entity"
)

type LogRepository interface {
	Insert(ctx context.Context, data entity.LogData) error
}
