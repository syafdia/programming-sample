package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/syafdia/sb/chal002/internal/domain/entity"

	mock_repo "github.com/syafdia/sb/chal002/internal/mock/domain/repo"
)

func Test_findMultipleMoviesUseCase_Execute(t *testing.T) {
	type fields struct {
		movieRepo *mock_repo.MockMovieRepository
		logRepo   *mock_repo.MockLogRepository
	}
	type args struct {
		ctx        context.Context
		searchWord string
		pagination int
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name       string
		fields     fields
		args       args
		beforeTest func(f fields)
		want       []entity.MovieSummary
		wantErr    bool
	}{
		{
			name: "Success finding multiple movies",
			fields: fields{
				logRepo:   mock_repo.NewMockLogRepository(ctrl),
				movieRepo: mock_repo.NewMockMovieRepository(ctrl),
			},
			args: args{
				ctx:        context.TODO(),
				searchWord: "Batman",
				pagination: 1,
			},
			beforeTest: func(f fields) {
				f.logRepo.
					EXPECT().
					Insert(
						context.TODO(),
						map[string]interface{}{
							"action": "MultipleMovies",
							"payload": map[string]interface{}{
								"search_word": "Batman",
								"pagination":  1,
							},
						},
					).
					Return(nil)

				f.movieRepo.
					EXPECT().
					FindMultipleSummaries(
						context.TODO(),
						"Batman",
						1,
					).Return(
					[]entity.MovieSummary{
						{
							Id:    "mv101010",
							Title: "Batman",
						},
					},
					nil,
				)
			},
			want: []entity.MovieSummary{
				{
					Id:    "mv101010",
					Title: "Batman",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &findMultipleMoviesUseCase{
				movieRepo: tt.fields.movieRepo,
				logRepo:   tt.fields.logRepo,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(tt.fields)
			}

			got, err := f.Execute(tt.args.ctx, tt.args.searchWord, tt.args.pagination)
			if (err != nil) != tt.wantErr {
				t.Errorf("findMultipleMoviesUseCase.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMultipleMoviesUseCase.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
