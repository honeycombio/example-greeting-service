require 'sinatra'
require 'honeycomb-beeline'
require 'faraday'

Honeycomb.configure do |config|
  config.write_key = ENV['HONEYCOMB_WRITE_KEY']
  config.dataset = ENV['HONEYCOMB_DATASET']
  config.service_name = "ruby-greeting-service-name"
  config.trace_http_header_propagation_hook do |request|

  end
  config.trace_http_header_parser_hook do |request|
    ["foo", "bar"]
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
  year_service_connection = Faraday.new("http://localhost:6000")
  year_service_response = year_service_connection.get("/year") do |request|
  end
  year_service_response.body.to_i
end
