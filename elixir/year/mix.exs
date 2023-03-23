defmodule Year.MixProject do
  use Mix.Project

  def project do
    [
      app: :year,
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
      mod: {Year.Application, []}
    ]
  end

  # Run "mix help deps" to learn about dependencies.
  defp deps do
    [
      {:plug_cowboy, "~> 2.6"},
      {:opentelemetry, "~> 1.2.0"},
      {:opentelemetry_api, "~> 1.2.1"},
      {:opentelemetry_exporter, "~> 1.4.0"},
      {:opentelemetry_process_propagator, "~> 0.2.2"},
      {:opentelemetry_cowboy, "~> 0.2"}
    ]
  end
end
