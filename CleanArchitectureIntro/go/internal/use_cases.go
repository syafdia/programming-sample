package internal

import "context"

type AppManagementUseCase interface {
	CreateApplication(ctx context.Context, req AppManagementCreateAppPayload) error
}

type AuthUseCase interface {
	HasAuthentication(ctx context.Context, req AuthenticationPayload) (bool, error)
	HasAuthorization(ctx context.Context, req AuthorizationPayload) (bool, error)
}

type DefaultAppManagementUseCase struct {
	aivenSvcRepo     AivenServiceRepository
	gcloudAppEngRepo GCloudAppEngineRepository
}

func NewAppManagementUseCase(
	aivenSvcRepo AivenServiceRepository,
	gcloudAppEngRepo GCloudAppEngineRepository,
) *DefaultAppManagementUseCase {
	return &DefaultAppManagementUseCase{
		aivenSvcRepo:     aivenSvcRepo,
		gcloudAppEngRepo: gcloudAppEngRepo,
	}
}

func (am *DefaultAppManagementUseCase) CreateApplication(ctx context.Context, req AppManagementCreateAppPayload) error {
	return nil
}

type DefaultAuthUseCase struct {
}

func NewAuthUseCase() *DefaultAuthUseCase {
	return &DefaultAuthUseCase{}
}

func (au *DefaultAuthUseCase) HasAuthentication(ctx context.Context, req AuthenticationPayload) (bool, error) {
	return false, nil
}

func (au *DefaultAuthUseCase) HasAuthorization(ctx context.Context, req AuthorizationPayload) (bool, error) {
	return false, nil
}
