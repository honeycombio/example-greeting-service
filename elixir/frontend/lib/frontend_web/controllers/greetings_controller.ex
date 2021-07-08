defmodule FrontendWeb.GreetingsController do
  use FrontendWeb, :controller

  require OpenTelemetry.Tracer, as: Tracer

  def index(conn, _params) do
    name = Frontend.name()
    message = Frontend.message()
    Tracer.with_span "🎨 render message ✨" do
      text(conn, "Hello #{name}, #{message}")
    end
  end
end
