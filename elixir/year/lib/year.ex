defmodule Year do
  use Plug.Router

  require OpenTelemetry.Tracer, as: Tracer

  plug :match
  plug :dispatch

  get "/year" do
    conn
    |> put_resp_content_type("text/plain")
    |> send_resp(200, year() |> to_string())
  end

  defp year do
    Tracer.with_span "🗓 get-a-year ✨" do
      :timer.sleep(Enum.random(0..5))
      Enum.random(2015..2020)
    end
  end
end
