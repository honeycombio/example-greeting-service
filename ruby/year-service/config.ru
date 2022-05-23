# frozen_string_literal: true

require "bundler/setup"
Bundler.require

require "honeycomb/propagation/w3c"

Honeycomb.configure do |config|
  config.write_key = ENV["HONEYCOMB_API_KEY"]
  config.service_name = "year-ruby"
  config.api_host = ENV['HONEYCOMB_API_ENDPOINT']

  config.http_trace_parser_hook do |env|
    Honeycomb::W3CPropagation::UnmarshalTraceContext.parse_rack_env(env)
  end

  config.notification_events = [/grape$/]
end

class App < Grape::API
  format :txt

  get "/year" do
    Honeycomb.start_span(name: "🗓 get-a-year ✨") do
      sleep rand(0..0.005)
      (2015..2020).to_a.sample
    end
  end
end

use Honeycomb::Rack::Middleware, client: Honeycomb.client
run App
