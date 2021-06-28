# frozen_string_literal: true

# SPDX-License-Identifier: Apache-2.0

require 'action_controller/railtie'
require 'faraday'

require 'opentelemetry/sdk'
require 'opentelemetry/exporter/otlp'
require 'opentelemetry/instrumentation/all'
begin
  OpenTelemetry::SDK.configure do |c|
    c.service_name = ENV['SERVICE_NAME'] || "frontend-ruby"
    c.use_all()
  end
rescue OpenTelemetry::SDK::ConfigurationError => e
  puts "What now?"
  puts e.inspect
end

Tracer = OpenTelemetry.tracer_provider.tracer('frontend-internal')

# Frontend is a minimal Rails application inspired by the Rails
# bug report template for action controller.
# The configuration is compatible with Rails 6.0
class FrontendGreetingApp < Rails::Application
  config.root = __dir__
  config.hosts << 'example.org'
  secrets.secret_key_base = 'secret_key_base'
  config.eager_load = false
  config.logger = Logger.new($stdout)
  Rails.logger  = config.logger

  routes.append do
    get "/greeting" => "greetings#index"
  end
end

class GreetingsController < ActionController::Base
  def index
    @name = NameClient.get_name
    @message = MessageClient.get_message
    Tracer.in_span("🎨 render message ✨") do |span|
      render inline: "Hello <%= @name %>, <%= @message %>"
    end
  end
end

class NameClient
  def self.get_name
    Tracer.in_span("✨ call /name ✨") do |_span|
      Faraday.new(ENV["NAME_ENDPOINT"] || "http://localhost:8000")
            .get("/name")
            .body
    end
  end
end

class MessageClient
  def self.get_message
    Tracer.in_span("✨ call /message ✨") do |_span|
      Faraday.new(ENV["MESSAGE_ENDPOINT"] || "http://localhost:9000")
            .get("/message")
            .body
    end
  end
end

Rails.application.initialize!

Rack::Server.new(app: FrontendGreetingApp, Port: 7000).start

# To run this example run the `rackup` command with this file
# Example: rackup frontend.ru
# Navigate to http://localhost:7000/
