import Config

config :opentelemetry, :resource,
  service: [
    name: "message-elixir"
  ]

config :opentelemetry,
  processors: [
    otel_batch_processor: %{
      exporter: {
        OpenTelemetry.Honeycomb.Exporter,
        write_key: System.get_env("HONEYCOMB_API_KEY"),
        dataset: System.get_env("HONEYCOMB_DATASET")
      }
    }
  ]
