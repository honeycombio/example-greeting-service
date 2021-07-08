import Config

config :opentelemetry, :resource, service: [name: "name-elixir"]

config :opentelemetry, :propagators, :otel_propagator_http_w3c

config :opentelemetry, :processors,
  otel_batch_processor: %{
    exporter: {
      :opentelemetry_exporter,
      %{
        endpoints: [
          {
            :http,
            System.get_env("OTEL_COLLECTOR_HOST", "localhost"),
            System.get_env("OTEL_COLLECTOR_PORT", "55681") |> String.to_integer(),
            []
          }
        ]
      }
    }
  }
