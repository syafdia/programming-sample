from ast import List
import logging
from python.entities import AppManagementCreateAppPayload

from python.use_cases import AppManagementUseCase


class CliHandler:
    def __init__(self, app_management_uc: AppManagementUseCase) -> None:
        self._app_management_uc = app_management_uc


    def create_application(self, args: List[str]) -> None:
        # TODO: Parse args from sys.argv to use case from CLI

        self._app_management_uc.create_application(AppManagementCreateAppPayload(
            # TODO: Construct params.
            git_uri=''
        ))