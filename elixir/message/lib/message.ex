defmodule Message do
  use Plug.Router

  require OpenTelemetry.Tracer, as: Tracer

  plug :match
  plug :dispatch

  @messages [
    "how are you?",
    "how are you doing?",
    "what's good?",
    "what's up?",
    "how do you do?",
    "sup?",
    "good day to you",
    "how are things?",
    "howzit?",
    "woohoo"
  ]

  get "/message" do
    :otel_propagator.text_map_extract(conn.req_headers)

    Tracer.with_span "/message" do
      Tracer.set_attributes(
        "http.method": conn.method,
        "http.scheme": conn.scheme,
        "http.host": conn.host,
        "http.route": conn.request_path
      )

      conn
      |> put_resp_content_type("text/plain")
      |> send_resp(200, message())
    end
  end

  defp message do
    Tracer.with_span "📖 look up message ✨" do
      :timer.sleep(Enum.random(0..5))
      Enum.random(@messages)
    end
  end
end
