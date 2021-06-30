defmodule FrontendWeb.Router do
  use FrontendWeb, :router

  get "/greeting", FrontendWeb.GreetingsController, :index
end
