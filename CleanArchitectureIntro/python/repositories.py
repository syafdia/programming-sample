from entities import (
    AivenCreateAuthenticationMethodPayload, 
    AivenAuthenticationMethod,
    SlackPostMessagePayload,
    AivenCreateServicePayload,
    AivenService,
)

from typing import Any, Dict, List, Tuple

class GcloudAppEngineRepository:
    pass

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

class TwilioRepository:
    def send_sms(self):
        pass