# frozen_string_literal: true

require 'logger'
require 'redis'

WORKER_ID = ARGV.length.positive? ? ARGV[0] : 0

logger = Logger.new($stdout)
redis_client = Redis.new(host: '0.0.0.0', port: 7379)

logger.info "Starting worker #{WORKER_ID}"

loop do
  _, email = redis_client.brpop('queue', 0)

  logger.info "Start sending welcome email to #{email}"
  sleep(5)
  logger.info "Welcome email has been sent to #{email}"
end
