require 'sinatra'
require 'honeycomb-beeline'
require 'honeycomb/propagation/w3c'
require 'faraday'

Honeycomb.configure do |config|
  config.write_key = ENV['HONEYCOMB_WRITE_KEY']
  config.dataset = ENV['HONEYCOMB_DATASET']
  config.service_name = "name-service-rb"
  config.http_trace_parser_hook do |env|
    Honeycomb::W3CPropagation::UnmarshalTraceContext.parse_rack_env(env)
  end
  config.http_trace_propagation_hook do |env, context|
    Honeycomb::W3CPropagation::MarshalTraceContext.parse_faraday_env(env, context)
  end
  #config.client = Libhoney::LogClient.new
end

use Honeycomb::Sinatra::Middleware, client: Honeycomb.client

set :port, 8000

names_by_year = {
  2015 => ['sophia', 'jackson', 'emma', 'aiden', 'olivia', 'liam', 'ava', 'lucas', 'mia', 'noah'],
  2016 => ['sophia', 'jackson', 'emma', 'aiden', 'olivia', 'lucas', 'ava', 'liam', 'mia', 'noah'],
  2017 => ['sophia', 'jackson', 'olivia', 'liam', 'emma', 'noah', 'ava', 'aiden', 'isabella', 'lucas'],
  2018 => ['sophia', 'jackson', 'olivia', 'liam', 'emma', 'noah', 'ava', 'aiden', 'isabella', 'caden'],
  2019 => ['sophia', 'liam', 'olivia', 'jackson', 'emma', 'noah', 'ava', 'aiden', 'aira', 'grayson'],
  2020 => ['olivia', 'noah', 'emma', 'liam', 'ava', 'elijah', 'isabella', 'oliver', 'sophia', 'lucas']
}

get '/name' do
  year = get_year
  names = names_by_year[year]
  names[rand(names.length)]
end

def get_year
  year_service_connection = Faraday.new("http://localhost:6001")
  year_service_response = year_service_connection.get("/year") do |request|
  end
  year_service_response.body.to_i
end
