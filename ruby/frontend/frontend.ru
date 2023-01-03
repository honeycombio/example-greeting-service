# frozen_string_literal: true

# SPDX-License-Identifier: Apache-2.0

require 'action_controller/railtie'
require 'faraday'

require 'opentelemetry/sdk'
require 'opentelemetry/exporter/otlp'
require 'opentelemetry/instrumentation/all'
require_relative './o11y_wrapper.rb'

begin
  OpenTelemetry::SDK.configure do |c|
    c.service_name = ENV['SERVICE_NAME'] || "frontend-ruby"

    # enable all auto-instrumentation available
    c.use_all()

    # add the Baggage and CarryOn processors to thepipeline
    c.add_span_processor(O11yWrapper::BaggageSpanProcessor.new)
    c.add_span_processor(O11yWrapper::CarryOnSpanProcessor.new)

    # Because we tinkered with the pipeline, we'll need to
    # wire up span batching and sending via OTLP ourselves.
    # This is usually the default.
    c.add_span_processor(
      OpenTelemetry::SDK::Trace::Export::BatchSpanProcessor.new(
        OpenTelemetry::Exporter::OTLP::Exporter.new()
      )
    )
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
    # set something in CarryOn that will appear in all child spans from the frontend,
    # but will NOT appear on child spans from other services
    O11yWrapper::CarryOn.with_attributes({"app.carry_on" => "my wayward son"}) do

      # a span event!
      OpenTelemetry::Trace
        .current_span
        .add_event("Emoji are fun! ᕕ( ᐛ )ᕗ")

      @name = NameClient.get_name

      # set name in Baggage for child spans, both frontend and other services from
      # this point in the trace forward
      OpenTelemetry::Context.with_current(OpenTelemetry::Baggage.set_value("app.visitor_name", @name)) do
        @message = MessageClient.get_message

        Tracer.in_span("🎨 render greeting ✨") do |span|
          render inline: "Hello <%= @name %>, <%= @message %>"
        end
      end
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

Rack::Server.new(app: FrontendGreetingApp, Port: 7007).start

# To run this example run the `rackup` command with this file
# Example: rackup frontend.ru
# Navigate to http://localhost:7007/
