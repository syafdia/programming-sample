# frozen_string_literal: true

require 'logger'
require 'redis'
require 'sinatra'
require 'sinatra/json'

logger = Logger.new($stdout)
redis_client = Redis.new(host: '0.0.0.0', port: 7379)

post '/send_welcome_email' do
  redis_client.lpush('queue', params[:email])
  logger.info "#{params[:email]} has been added to queue"

  status 201
  json message: 'Accepted'
end
