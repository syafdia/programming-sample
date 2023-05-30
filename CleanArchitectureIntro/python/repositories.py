from entities import (
    AivenCreateAuthenticationMethodPayload, 
    AivenAuthenticationMethod,
    SlackPostMessagePayload,
    AivenCreateServicePayload,
    AivenService,
)

from typing import Any, Dict, List, Tuple

from python.entities import GcloudAppEngineCreateAppPayload, GcloudAppEngineDeployAppPayload, GithubApplicationMetadata

class GcloudAppEngineRepository:

    def __init__(self, gcloud_client: Any, bash: Any) -> None:
        self._gcloud_client = gcloud_client
        self._bash = bash


    def create_application(self, payload: GcloudAppEngineCreateAppPayload) -> None:
        """
        Call appengine_admin_v1.ApplicationsClient.create_application 
        based on https://cloud.google.com/python/docs/reference/appengine/latest/google.cloud.appengine_admin_v1.services.applications.ApplicationsClient#google_cloud_appengine_admin_v1_services_applications_ApplicationsClient_create_application
        """
        
        self._gcloud_client.create_application(payload)

    def deploy_application(self, payload: GcloudAppEngineDeployAppPayload) -> None:
        """
        Run `gcloud app deploy` command
        """

        self._bash([
            'gcloud', 'app', 'deploy', f'--appyaml={payload.app_yaml}'
        ])

    def delete_application(self, service_name: str) -> None:
        """
        Run `gcloud app service delete {service_name}`
        """

        self._bash([
            'gcloud', 'app', 'service', 'delete', service_name
        ])

class AivenAccountRepository:
    """
    Represents operation on https://api.aiven.io/doc/#tag/Account
    """

    def create_authentication_method(
            self, 
            account_id: str, 
            payload: AivenCreateAuthenticationMethodPayload,
            ) -> AivenAuthenticationMethod:
        """
        Represents POST https://api.aiven.io/v1/account/{account_id}/authentication
        """
        pass 

    def list_authentication_method(
            self, 
            account_id: str,
            ) -> List[AivenAuthenticationMethod]:
        """
        Represents GET https://api.aiven.io/v1/account/{account_id}/authentication
        """
        pass

    def delete_authentication_method(
            self, 
            account_id: str,
            authentication_id: str,
        ) -> None:
        """
        Represents aiven-cli via `avn account authentication-method delete {account_id} {authentication_id}`
        """

        pass

class AivenServiceRepository:
    def __init__(self, aiven_client: None) -> None:
        self._aiven_client = aiven_client

    def create_service(self, project: str, payload: AivenCreateServicePayload) -> AivenService: 
        """
        Represents POST https://api.aiven.io/v1/project/{project}/service
        """
        pass

class SlackRepository:

    def post_mesage(self, payload: SlackPostMessagePayload) -> None:
        """
        Represents POST https://slack.com/api/chat.postMessage
        """
        pass 

class GithubRepository:
    def get_application_metadata(self, git_uri) -> GithubApplicationMetadata:
        pass

class SendGridRepository:
    def send_mail(self, to_addresses: List[str], body: str) -> None:
        pass