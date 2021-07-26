# frozen_string_literal: true
# rubocop:disable Style/GlobalVars

require 'redis'
require 'sinatra'
require 'sinatra/json'

POKEMONS = [
  {
    'id': 1,
    'name': 'Bulbasaur'
  },
  {
    'id': 2,
    'name': 'Ivysaur'
  },
  {
    'id': 3,
    'name': 'Venusaur'
  },
  {
    'id': 4,
    'name': 'Charmander'
  },
  {
    'id': 5,
    'name': 'Charmeleon'
  },
  {
    'id': 6,
    'name': 'Charizard'
  },
  {
    'id': 7,
    'name': 'Squirtle'
  },
  {
    'id': 8,
    'name': 'Wartortle'
  },
  {
    'id': 9,
    'name': 'Blastoise'
  }
].freeze

$redis_client = Redis.new(host: '0.0.0.0', port: 7379)

get '/pokemons' do
  limiter_key = request.ip
  tot_request = $redis_client.get(limiter_key).to_i

  # Only 5 requests per minute
  if tot_request > 5
    status 429
    return json message: 'Too many request'
  end

  $redis_client.incr(limiter_key)
  $redis_client.expire(limiter_key, 59)

  json POKEMONS
end
