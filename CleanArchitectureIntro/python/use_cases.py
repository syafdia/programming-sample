from repositories import (
    GcloudAppEngineRepository,
    AivenServiceRepository
)

from entities import (
    AppManagementCreateAppPayload
)

class AppManagementUseCase:
    def __init__(
            self, 
            gcloud_app_engine_repo: GcloudAppEngineRepository,
            aiven_service_repo: AivenServiceRepository,
        ) -> None:
        
        self.gcloud_app_engine_repo = gcloud_app_engine_repo
        self.aiven_service_repo = aiven_service_repo

    def create_application(self, req: AppManagementCreateAppPayload):
        pass

    def restart_application(self):
        pass

    def destroy_application(self):
        pass