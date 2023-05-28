from dataclasses import dataclass
from typing import Union
import datetime

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
class AppManagementCreateAppPayload:
    git_uri: str