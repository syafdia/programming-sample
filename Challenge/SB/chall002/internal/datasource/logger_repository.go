package datasource

import (
	"context"
	"encoding/json"
	"log"

	"github.com/syafdia/sb/chal002/internal/domain/entity"
	"github.com/syafdia/sb/chal002/internal/domain/repo"
)

type logRepo struct {
}

func NewLogRepo() repo.LogRepository {
	return &logRepo{}
}

func (l *logRepo) Insert(ctx context.Context, data entity.LogData) error {
	go func() {
		jsonData, _ := json.Marshal(&data)
		log.Printf("logRepo: Saving log data to DB, data: %s\n", string(jsonData))
	}()
	return nil
}
