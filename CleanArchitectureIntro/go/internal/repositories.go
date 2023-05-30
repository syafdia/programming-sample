package internal

import (
	"context"

	"github.com/syafdia/programming-sample/CleanArchitectureIntro/pkg/types"
)

// List of repository contracts.

type GCloudAppEngineRepository interface {
	DeployApplication(ctx context.Context, payload GCloudAppEngineDeployAppPayload) error
}

type AivenServiceRepository interface {
	CreateService(ctx context.Context, project string, payload AivenCreateServicePayload) (AivenService, error)
}

// DefaultAivenServiceRepository implements AivenServiceRepository interface.
type DefaultAivenServiceRepository struct {
	aivenClient types.AivenClient
	bash        types.BashExecutor
}

func NewAivenServiceRepository(
	aivenClient types.AivenClient,
	bash types.BashExecutor,
) *DefaultAivenServiceRepository {
	return &DefaultAivenServiceRepository{
		aivenClient: aivenClient,
		bash:        bash,
	}
}

func (as *DefaultAivenServiceRepository) CreateService(ctx context.Context, project string, payload AivenCreateServicePayload) (AivenService, error) {
	return AivenService{}, nil
}

// DefaultGCloudAppEngineRepository implements GCloudAppEngineRepository interface.
type DefaultGCloudAppEngineRepository struct {
	gcloudClient types.GCloudClient
	bash         types.BashExecutor
}

func NewGCloudAppEngineRepository(
	gcloudClient types.GCloudClient,
	bash types.BashExecutor,
) *DefaultGCloudAppEngineRepository {
	return &DefaultGCloudAppEngineRepository{
		gcloudClient: gcloudClient,
		bash:         bash,
	}
}

// deploy_application(self, payload: GCloudAppEngineDeployAppPayload) -> None:
func (ae *DefaultGCloudAppEngineRepository) DeployApplication(ctx context.Context, payload GCloudAppEngineDeployAppPayload) error {
	return nil
}
