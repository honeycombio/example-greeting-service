# frozen_string_literal: true

require "bundler/setup"
Bundler.require

require 'opentelemetry/sdk'
require 'opentelemetry/exporter/otlp'
require 'opentelemetry/instrumentation/all'
require_relative './o11y_wrapper.rb'

begin
  OpenTelemetry::SDK.configure do |c|
    c.service_name = ENV['SERVICE_NAME'] || "year-ruby"

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

Tracer = OpenTelemetry.tracer_provider.tracer("year-internal")

class App < Grape::API
  format :txt

  get :year do
    Tracer.in_span("🗓 get-a-year ✨") do
      sleep rand(0..0.005)
      (2015..2020).to_a.sample
    end
  end
end

use OpenTelemetry::Instrumentation::Rack::Middlewares::TracerMiddleware
run App
