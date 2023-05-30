package main

import (
	"log"
	"net/http"

	"github.com/syafdia/programming-sample/CleanArchitectureIntro/internal"
	"github.com/syafdia/programming-sample/CleanArchitectureIntro/pkg/types"
)

func main() {
	log.Println("[main] OK")

	// Init libraries.
	var gcloudClient types.GCloudClient
	var aivenClient types.AivenClient
	var bash types.BashExecutor

	// Init repositories layer.
	aivenServiceRepo := internal.NewAivenServiceRepository(aivenClient, bash)
	gcloudAppEngineRepo := internal.NewGCloudAppEngineRepository(gcloudClient, bash)

	// Init use cases layer.
	appManagementUseCase := internal.NewAppManagementUseCase(aivenServiceRepo, gcloudAppEngineRepo)
	authUseCase := internal.NewAuthUseCase()

	// Init deliveries layer as HTTP server.
	handler := internal.NewHTTPHandler(appManagementUseCase, authUseCase)
	http.HandleFunc("/app_management", handler.DestroyApplication)

	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Panicf("[main] failed when run HTTP server, error: %v", err)
	}

}
