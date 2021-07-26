# frozen_string_literal: true
# rubocop:disable Style/GlobalVars

require 'json'
require 'logger'
require 'redis'

$logger = Logger.new($stdout)
$redis_pub = Redis.new(host: '0.0.0.0', port: 7379)
$redis_sub = Redis.new(host: '0.0.0.0', port: 7379)

def make_payment(payload)
  $logger.info "Make payment for #{payload}"
  sleep(3)

  if (payload['user_id'].to_i & 1) == 1
    $logger.info "Payment for #{payload} has been succeed"
    $redis_pub.publish(:MAKE_PAYMENT_SUCCESS, payload.to_json)
    return
  end

  $logger.info "Payment for #{payload} failed"
  $redis_pub.publish(:MAKE_PAYMENT_FAILED, payload.to_json)
end

def refund(payload)
  $logger.info "Refund payment for #{payload}"
  sleep(3)
  $logger.info "Refund for #{payload} has been succeed"
end

$logger.info 'Starting payment pubsub...'
$redis_sub.subscribe(:BLOCK_SEAT_SUCCESS, :ALLOCATE_SEAT_FAILED) do |on|
  on.message do |channel, message|
    payload = JSON.parse(message)

    case channel.to_sym
    when :BLOCK_SEAT_SUCCESS
      make_payment(payload)
    when :ALLOCATE_SEAT_FAILED
      refund(payload)
    else
      $logger.info "Unhandled channel #{channel}"
    end
  end
end
