# This file is responsible for configuring your application
# and its dependencies with the aid of the Mix.Config module.
#
# This configuration file is loaded before any dependency and
# is restricted to this project.

# General application configuration
import Config

# Configures the endpoint
config :frontend, FrontendWeb.Endpoint,
  url: [host: "localhost"],
  secret_key_base: "28gPO3Atbcv0ZX9ATXD6/prb3MOwl6DHGGbHWP8jEIzP4xV4/GPftCS+BtGqgO7n",
  render_errors: [view: FrontendWeb.ErrorView, accepts: ~w(json), layout: false],
  pubsub_server: Frontend.PubSub,
  live_view: [signing_salt: "RCSHpWzo"]

# Configures Elixir's Logger
config :logger, :console,
  format: "$time $metadata[$level] $message\n",
  metadata: [:request_id]

# Use Jason for JSON parsing in Phoenix
config :phoenix, :json_library, Jason

# Configures OpenTelemetry
config :opentelemetry, :resource, service: [name: "frontend-elixir"]
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

# Import environment specific config. This must remain at the bottom
# of this file so it overrides the configuration defined above.
import_config "#{Mix.env()}.exs"
