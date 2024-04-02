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

    # enable all instrumentation libraries installed in the Gemfile
    # and configure some
    c.use_all({
      'OpenTelemetry::Instrumentation::Rack' => {
        untraced_endpoints: ['/up'], # don't create traces for healthchecks
      },
    })

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
    get "/up" => "health#show"
    get "/greeting" => "greetings#index"
  end
end

class HealthController < ActionController::Base
  def show
    render plain: "ok"
  end
end

class GreetingsController < ActionController::Base
  around_action :context_for_named_visitor

  def index
    # a span event!
    OpenTelemetry::Trace
      .current_span
      .add_event("GreetingsController#index starts it's specific behavior! ᕕ( ᐛ )ᕗ")


    # set something in CarryOn that will appear in all child spans from the frontend,
    # but will NOT appear on child spans from other services
    O11yWrapper::CarryOn.with_attributes({"app.carry_on" => "my wayward son"}) do


      @message = MessageClient.get_message

      Tracer.in_span("🎨 render greeting ✨") do |span|
        render inline: "Hello <%= @name %>, <%= @message %>"
      end
    end
  end

  private

  # an around filter to retrieve a name for the greeting and
  # store that name (once known) in Baggage
  def context_for_named_visitor
    # get a name from the Name service
    @name = NameClient.get_name

    # add the name as an attribute to the current span
    OpenTelemetry::Trace
      .current_span
      .set_attribute("app.visitor_name", @name)

    # store that name in Baggage (which is in Context) to be applied by
    # the BaggageSpanProcessor as an attribute on child spans
    named_visitor_context = OpenTelemetry::Baggage.set_value("app.visitor_name", @name)

    # set that Context as current for the activity within a controller action
    OpenTelemetry::Context.with_current(named_visitor_context) do
      yield
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

Rack::Server.new(app: FrontendGreetingApp, Port: 7777).start

# To run this example run the `rackup` command with this file
# Example: rackup frontend.ru
# Navigate to http://localhost:7777/
