import logging
import sys
from python.deliveries import CliHandler

from python.repositories import AivenServiceRepository, GcloudAppEngineRepository, GithubRepository, SendGridRepository, SlackRepository
from python.use_cases import AppManagementUseCase

def main():
  logging.info('[main] Run CLI app')

  # Init libraries.
  aiven_client = None
  gcloud_client = None
  bash = None

  # Init repositories layer.
  gcloud_app_engine_repo = GcloudAppEngineRepository(gcloud_client, bash)
  aiven_service_repo =  AivenServiceRepository(aiven_client)
  github_repo =  GithubRepository()
  slack_repo =  SlackRepository()
  sendgrid_repo =  SendGridRepository()

  # Init use cases layer.
  app_management_uc = AppManagementUseCase(
    gcloud_app_engine_repo=gcloud_app_engine_repo,
    aiven_service_repo=aiven_service_repo,
    github_repo=github_repo,
    slack_repo=slack_repo,
    sendgrid_repo=sendgrid_repo,
  )


  # Init deliveries layer as CLI application
  cli_handler = CliHandler(app_management_uc=app_management_uc)

  action = sys.argv[0]
  args = sys.argv[1:]

  if action == 'create_application':
    cli_handler.create_application(args)
  elif action == 'describe_application':
    # TODO: Call describe_application method
    pass
  else:
    raise Exception(f'Action ${action} is not supported')



if __name__ == '__main__':
  main()