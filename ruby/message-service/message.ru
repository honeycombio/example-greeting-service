# frozen_string_literal: true

# SPDX-License-Identifier: Apache-2.0

require 'action_controller/railtie'
require 'honeycomb-beeline'
require 'honeycomb/propagation/w3c'
require 'faraday'

Honeycomb.configure do |config|
  config.write_key = ENV['HONEYCOMB_API_KEY']
  config.dataset = ENV['HONEYCOMB_DATASET']
  config.service_name = ENV['SERVICE_NAME'] || "frontend-ruby"

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

# MessageApp is a minimal Rails application inspired by the Rails
# bug report template for action controller.
# The configuration is compatible with Rails 6.0
class MessageApp < Rails::Application
  config.root = __dir__
  config.hosts << 'example.org'
  secrets.secret_key_base = 'secret_key_base'
  config.eager_load = false
  config.logger = Logger.new($stdout)
  Rails.logger  = config.logger

  routes.append do
    get "/message" => "messages#index"
  end
end


class MessagesController < ActionController::Base
  MESSAGES = [
    "how are you?", "how are you doing?", "what's good?", "what's up?", "how do you do?",
    "sup?", "good day to you", "how are things?", "howzit?", "woohoo",
  ]

  def index
    Honeycomb.add_field('name', '/message') # why doesn't this overwrite the event name?
    message = ""
    Honeycomb.start_span(name: "📖 look up message ✨") do |span|
      message = MESSAGES.sample
    end
    render plain: message
  end
end

Rails.application.initialize!

Rack::Server.new(app: MessageApp, Port: 9000).start

# To run this example run the `rackup` command with this file
# Example: rackup message.ru
# Navigate to http://localhost:9000/
