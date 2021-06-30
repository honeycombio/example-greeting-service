defmodule FrontendWeb.Endpoint do
  use Phoenix.Endpoint, otp_app: :frontend

  plug Plug.RequestId
  plug Plug.Telemetry, event_prefix: [:phoenix, :endpoint]
  plug FrontendWeb.Router
end
