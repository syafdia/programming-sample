package datasource

import (
	"context"
	"testing"

	"github.com/syafdia/sb/chal002/internal/domain/entity"
)

func Test_logRepo_Insert(t *testing.T) {
	type args struct {
		ctx  context.Context
		data entity.LogData
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success inserting log to DB",
			args: args{
				ctx: context.TODO(),
				data: map[string]interface{}{
					"name": "John",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &logRepo{}
			if err := l.Insert(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("logRepo.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
