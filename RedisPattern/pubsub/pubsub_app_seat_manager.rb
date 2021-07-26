# frozen_string_literal: true
# rubocop:disable Style/GlobalVars

require 'json'
require 'logger'
require 'redis'

$logger = Logger.new($stdout)
$redis_pub = Redis.new(host: '0.0.0.0', port: 7379)
$redis_sub = Redis.new(host: '0.0.0.0', port: 7379)

def unblock_seat(payload)
  $logger.info "Unblock seat for #{payload} ..."
  sleep(3)
  $logger.info "Seat for #{payload} has been unblocked"
end

def allocate_seat(payload)
  $logger.info "Allocate seat for #{payload} ..."
  sleep(3)

  if payload['seat_id'] != 'F-01'
    $logger.info "Seat for #{payload} has been allocated"
    $redis_pub.publish(:ALLOCATE_SEAT_SUCCESS, payload.to_json)
    return
  end

  $logger.info "Seat for #{payload} is not allocated"
  $redis_pub.publish(:ALLOCATE_SEAT_FAILED, payload.to_json)
end

$logger.info 'Starting seat manager pubsub...'
$redis_sub.subscribe(:MAKE_PAYMENT_SUCCESS, :MAKE_PAYMENT_FAILED, :ALLOCATE_SEAT_FAILED) do |on|
  on.message do |channel, message|
    payload = JSON.parse(message)

    case channel.to_sym
    when :MAKE_PAYMENT_SUCCESS
      allocate_seat(payload)
    when :MAKE_PAYMENT_FAILED, :ALLOCATE_SEAT_FAILED
      unblock_seat(payload)
    else
      $logger.info "Unhandled channel #{channel}"
    end
  end
end
