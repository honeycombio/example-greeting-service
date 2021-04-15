# frozen_string_literal: true

# SPDX-License-Identifier: Apache-2.0

require 'action_controller/railtie'
require 'honeycomb-beeline'
require 'honeycomb/propagation/w3c'
require 'faraday'

Honeycomb.configure do |config|
  config.write_key = ENV['HONEYCOMB_WRITE_KEY']
  config.dataset = ENV['HONEYCOMB_DATASET']
  config.service_name = "frontend-rails"

  config.http_trace_parser_hook do |env|
    Honeycomb::W3CPropagation::UnmarshalTraceContext.parse_rack_env(env)
  end
  config.http_trace_propagation_hook do |env, context|
    Honeycomb::W3CPropagation::MarshalTraceContext.parse_faraday_env(env, context)
  end

  config.notification_events = %w[
    sql.active_record
    render_template.action_view
    render_partial.action_view
    render_collection.action_view
    process_action.action_controller
    send_file.action_controller
    send_data.action_controller
    deliver.action_mailer
  ].freeze
end

# YearApp is a minimal Rails application inspired by the Rails
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
    Honeycomb.add_field('name', '/greeting') # why doesn't this overwrite the event name?
    name = NameClient.get_name
    message = MessageClient.get_message
    Honeycomb.start_span(name: "🎨 render message ✨") do |span|
      render plain: "Hello #{name}, #{message}"
    end
  end
end

class NameClient
  def self.get_name
    Honeycomb.start_span(name: "✨ call /name ✨") do |_span|
      Faraday.new("http://localhost:8000")
            .get("/name")
            .body
    end
  end
end

class MessageClient
  def self.get_message
    Honeycomb.start_span(name: "✨ call /message ✨") do |_span|
      Faraday.new("http://localhost:9000")
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
