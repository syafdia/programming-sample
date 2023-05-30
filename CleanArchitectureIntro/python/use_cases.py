from python.entities import SLACK_CHANNEL_ID_APP_MANAGEMENT, AivenCreateServicePayload, AppManagementDestroyAppPayload, GcloudAppEngineCreateAppPayload, GcloudAppEngineDeployAppPayload, SlackPostMessagePayload
from python.repositories import GithubRepository, SendGridRepository, SlackRepository
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
            github_repo: GithubRepository,
            slack_repo: SlackRepository,
            sendgrid_repo: SendGridRepository,
        ) -> None:
        
        self._github_repo = github_repo
        self._gcloud_app_engine_repo = gcloud_app_engine_repo
        self._aiven_service_repo = aiven_service_repo
        self._slack_repo = slack_repo
        self._sendgrid_repo = sendgrid_repo

    def create_application(self, req: AppManagementCreateAppPayload) -> None:
        app_metadata = self._github_repo.get_application_metadata(req.git_uri)

        database = self._aiven_service_repo.create_service(
            req.project_id,
            AivenCreateServicePayload(
                plan=app_metadata.aiven_database_plan,
                service_name=app_metadata.application_name,
                service_type=app_metadata.aiven_database_type
            )
        )

        self._slack_repo.post_mesage(SlackPostMessagePayload(
            channel=SLACK_CHANNEL_ID_APP_MANAGEMENT,
            text=f'Success create database for {app_metadata.application_name} with ID {database.id}'
        ))

        self._gcloud_app_engine_repo.deploy_application(GcloudAppEngineDeployAppPayload(
            app_yaml=f"f{req.git_uri}/app.yaml"
        ))

        self._slack_repo.post_mesage(SlackPostMessagePayload(
            channel=SLACK_CHANNEL_ID_APP_MANAGEMENT,
            text=f'Success deploy new application database {app_metadata.application_name}'
        ))

    def destroy_application(self, req: AppManagementDestroyAppPayload) -> None:
        app_metadata = self._github_repo.get_application_metadata(req.git_uri)
        
        self._gcloud_app_engine_repo.delete_application(app_metadata.application_name)

        self._slack_repo.post_mesage(SlackPostMessagePayload(
            channel=SLACK_CHANNEL_ID_APP_MANAGEMENT,
            text=f'Application {app_metadata.application_name} has been destroyed'
        ))

        self._sendgrid_repo.send_mail(req.notify_email_addresses, 'Application {app_metadata.application_name} has been destroyed')