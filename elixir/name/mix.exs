defmodule Name.MixProject do
  use Mix.Project

  def project do
    [
      app: :name,
      version: "0.1.0",
      elixir: "~> 1.12",
      start_permanent: Mix.env() == :prod,
      deps: deps()
    ]
  end

  # Run "mix help compile.app" to learn about applications.
  def application do
    [
      extra_applications: [:logger],
      mod: {Name.Application, []}
    ]
  end

  # Run "mix help deps" to learn about dependencies.
  defp deps do
    [
      {:plug_cowboy, "~> 2.5"},
      {:opentelemetry, "~> 1.0.0-rc"},
      {:opentelemetry_api, "~> 1.0.0-rc"},
      {:opentelemetry_exporter, "~> 1.0.0-rc"},
      {:opentelemetry_plug, github: "opentelemetry-beam/opentelemetry_plug"},
      {:httpoison, "~> 2.0"}
    ]
  end
end
