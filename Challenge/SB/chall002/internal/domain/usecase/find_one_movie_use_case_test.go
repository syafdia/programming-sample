package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/syafdia/sb/chal002/internal/domain/entity"
	mock_repo "github.com/syafdia/sb/chal002/internal/mock/domain/repo"
)

func Test_findOneMovieUseCase_Execute(t *testing.T) {
	type fields struct {
		movieRepo *mock_repo.MockMovieRepository
		logRepo   *mock_repo.MockLogRepository
	}
	type args struct {
		ctx context.Context
		id  string
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tests := []struct {
		name       string
		fields     fields
		args       args
		beforeTest func(f fields)
		want       entity.MovieDetail
		wantErr    bool
	}{
		{
			name: "Success finding one movie",
			fields: fields{
				logRepo:   mock_repo.NewMockLogRepository(ctrl),
				movieRepo: mock_repo.NewMockMovieRepository(ctrl),
			},
			args: args{
				ctx: context.TODO(),
				id:  "mv0001",
			},
			beforeTest: func(f fields) {
				f.logRepo.
					EXPECT().
					Insert(
						context.TODO(),
						map[string]interface{}{
							"action": "FindOneMovie",
							"payload": map[string]interface{}{
								"id": "mv0001",
							},
						},
					).
					Return(nil)

				f.movieRepo.
					EXPECT().
					FindOneDetail(
						context.TODO(),
						"mv0001",
					).Return(
					entity.MovieDetail{
						Id:    "mv0001",
						Title: "The Omen",
					},
					nil,
				)
			},
			want: entity.MovieDetail{
				Id:    "mv0001",
				Title: "The Omen",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &findOneMovieUseCase{
				movieRepo: tt.fields.movieRepo,
				logRepo:   tt.fields.logRepo,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(tt.fields)
			}

			got, err := f.Execute(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("findOneMovieUseCase.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOneMovieUseCase.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
