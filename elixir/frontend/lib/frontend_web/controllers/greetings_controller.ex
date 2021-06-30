defmodule FrontendWeb.GreetingsController do
  use FrontendWeb, :controller

  def index(conn, _params) do
    name = Frontend.name()
    message = Frontend.message()
    text(conn, "Hello #{name}, #{message}")
  end
end
