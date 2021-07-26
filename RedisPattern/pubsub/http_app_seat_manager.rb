# frozen_string_literal: true

require 'logger'
require 'redis'
require 'sinatra'
require 'sinatra/json'

logger = Logger.new($stdout)
redis_client = Redis.new(host: '0.0.0.0', port: 7379)

post '/block_seat' do
  payload = { event_id: Time.now.to_i, user_id: params[:user_id], seat_id: params[:seat_id] }

  logger.info "Start blocking seat #{payload}"
  sleep(3)
  redis_client.publish(:BLOCK_SEAT_SUCCESS, payload.to_json)
  logger.info "Seat #{payload[:seat_id]} has been bocked"

  json message: 'OK'
end
