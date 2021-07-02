require 'sinatra'
require 'honeycomb-beeline'
require 'honeycomb/propagation/w3c'
require 'faraday'

Honeycomb.configure do |config|
  config.write_key = ENV['HONEYCOMB_API_KEY']
  config.dataset = ENV['HONEYCOMB_DATASET']
  config.service_name = ENV['SERVICE_NAME'] || "name-ruby"
  config.http_trace_parser_hook do |env|
    Honeycomb::W3CPropagation::UnmarshalTraceContext.parse_rack_env(env)
  end
  config.http_trace_propagation_hook do |env, context|
    Honeycomb::W3CPropagation::MarshalTraceContext.parse_faraday_env(env, context)
  end
  #config.client = Libhoney::LogClient.new
end

use Honeycomb::Sinatra::Middleware, client: Honeycomb.client

set :bind, '0.0.0.0'
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
  Honeycomb.start_span(name: "📖 look up name based on year ✨") do |span|
    year = get_year
    span.add_field("app.year", year)
    name = names_by_year.fetch(year, ["OH NO!"]).sample
    span.add_field("app.username", name)
    name
  end
end

def get_year
  Honeycomb.start_span(name: "✨ call /year ✨") do |span|
    response = Faraday.new(ENV["YEAR_ENDPOINT"] || "http://localhost:6001").get("/year")
    year = response.body.to_i
    span.add_field("app.year", year)
    year
  end
end
