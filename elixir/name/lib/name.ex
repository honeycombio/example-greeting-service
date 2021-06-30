defmodule Name do
  use Plug.Router

  require OpenTelemetry.Tracer, as: Tracer

  plug :match
  plug :dispatch

  @names %{
    2015 => ~w[sophia jackson emma aiden olivia liam ava lucas mia noah],
    2016 => ~w[sophia jackson emma aiden olivia lucas ava liam mia noah],
    2017 => ~w[sophia jackson olivia liam emma noah ava aiden isabella lucas],
    2018 => ~w[sophia jackson olivia liam emma noah ava aiden isabella caden],
    2019 => ~w[sophia liam olivia jackson emma noah ava aiden aira grayson],
    2020 => ~w[olivia noah emma liam ava elijah isabella oliver sophia lucas]
  }

  get "/name" do
    conn
    |> put_resp_content_type("text/plain")
    |> send_resp(200, name_by(year()))
  end

  defp name_by(year) do
    Tracer.with_span "📖 look up name based on year ✨" do
      :timer.sleep(Enum.random(1..5))
      Map.get(@names, year) |> Enum.random()
    end
  end

  defp year do
    Tracer.with_span "✨ call /year ✨" do
      endpoint = System.get_env("YEAR_ENDPOINT", "http://localhost:6001")
      HTTPoison.get!("#{endpoint}/year").body |> String.to_integer()
    end
  end
end
