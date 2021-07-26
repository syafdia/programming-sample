# frozen_string_literal: true
# rubocop:disable Style/GlobalVars

require 'logger'
require 'redis'
require 'sinatra'
require 'sinatra/json'

$redis_client = Redis.new(host: '0.0.0.0', port: 7379)
$logger = Logger.new($stdout)

post '/generate_reports' do
  lock_key = request.ip
  lock_value = generate_random_str

  lock(lock_key, lock_value)

  generate_reports

  unlock(lock_key, lock_value)

  json message: 'OK'
rescue FailedAcquireLock => e
  $logger.info 'Other process failed acquiring lock'
  status 400
  json message: e.message
end

class FailedAcquireLock < StandardError; end

def lock(key, val)
  ok = $redis_client.set(key, val, nx: true, ex: 60)
  raise FailedAcquireLock, "Can not acquire lock with key #{key}" unless ok
end

def unlock(key, val)
  script = '
    if redis.call("GET",KEYS[1]) == ARGV[1]
    then
        return redis.call("DEL",KEYS[1])
    else
        return 0
    end
  '

  $redis_client.eval(script, keys: [key], argv: [val])
end

def generate_random_str
  Digest::MD5.hexdigest(Time.now.to_s)
end

def generate_reports
  $logger.info 'Start generating report...'
  sleep(15)
  $logger.info 'Report has been generated'
end
