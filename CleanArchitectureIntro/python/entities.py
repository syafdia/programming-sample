from dataclasses import dataclass
from typing import List, Union
import datetime

SLACK_CHANNEL_ID_APP_MANAGEMENT = 'XX-XXXXX'
SLACK_CHANNEL_ID_ERROR = 'YY-YYYYY'

@dataclass
class GcloudAppEngineCreateAppPayload:
    service_account: str
    region: str

@dataclass
class GcloudAppEngineDeployAppPayload:
    app_yaml: str

@dataclass
class AivenCreateAuthenticationMethodPayload:
    authentication_method_name: str
    authentication_method_type: str

@dataclass
class AivenAuthenticationMethod:
    account_id: str
    authentication_method_name: str
    authentication_method_type: str

@dataclass
class AivenCreateServicePayload:
    plan: str
    service_name: str
    service_type: str

@dataclass
class AivenService:
    id: str
    plan: str
    service_name: str
    service_type: str
    cloud_name: str
    create_time: datetime.datetime

@dataclass
class SlackPostMessagePayload:
    channel: str
    text: str
    thread_ts: Union[str, None] = None

@dataclass
class GithubApplicationMetadata:
    region: str
    application_name: str
    gcloud_project_id: str
    gcloud_service_account: str
    aiven_database_plan: str
    aiven_database_type: str

@dataclass
class AppManagementCreateAppPayload:
    git_uri: str

class AppManagementDestroyAppPayload:
    git_uri: str
    notify_email_addresses: List[str]

