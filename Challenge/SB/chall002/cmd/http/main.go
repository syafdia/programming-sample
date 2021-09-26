package main

import (
	"github.com/gin-gonic/gin"
	httpdelivery "github.com/syafdia/sb/chal002/internal/delivery/http"
	"github.com/syafdia/sb/chal002/internal/di"
)

func main() {
	// Setup module.
	appModule := di.NewAppModule()
	repoModule := di.NewRepoModule(appModule)
	useCaseModule := di.NewUseCaseModule(repoModule)

	// Setup HTTP handler.
	movieHandler := httpdelivery.NewMovieHandler(
		useCaseModule.FindOneMovieUseCase,
		useCaseModule.FindMultipleMoviesUseCase,
	)

	// Setup HTTP Server.
	r := gin.Default()
	r.GET("/movies", movieHandler.FindMultiple)
	r.GET("/movies/:id", movieHandler.FindOne)
	r.Run(":8001")
}
