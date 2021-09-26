package di

import (
	"net/http"
	"time"

	"github.com/syafdia/sb/chal002/internal/datasource"
	"github.com/syafdia/sb/chal002/internal/domain/repo"
	"github.com/syafdia/sb/chal002/internal/domain/usecase"
)

type AppModule struct {
	HTTPClient *http.Client
}

type RepoModule struct {
	LogRepository   repo.LogRepository
	MovieRepository repo.MovieRepository
}

type UseCaseModule struct {
	FindOneMovieUseCase       usecase.FindOneMovieUseCase
	FindMultipleMoviesUseCase usecase.FindMultipleMoviesUseCase
}

func NewAppModule() *AppModule {
	return &AppModule{
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func NewRepoModule(appModule *AppModule) *RepoModule {
	return &RepoModule{
		LogRepository:   datasource.NewLogRepo(),
		MovieRepository: datasource.NewMovieRepo(appModule.HTTPClient),
	}
}

func NewUseCaseModule(repoModule *RepoModule) *UseCaseModule {
	return &UseCaseModule{
		FindOneMovieUseCase:       usecase.NewFindOneMovieUseCase(repoModule.MovieRepository, repoModule.LogRepository),
		FindMultipleMoviesUseCase: usecase.NewFindMultipleMoviesUseCase(repoModule.MovieRepository, repoModule.LogRepository),
	}
}
