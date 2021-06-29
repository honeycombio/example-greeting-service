defmodule Message.MixProject do
  use Mix.Project

  def project do
    [
      app: :message,
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
      mod: {Message.Application, []}
    ]
  end

  # Run "mix help deps" to learn about dependencies.
  defp deps do
    [
      {:plug_cowboy, "~> 2.5"},
      {:opentelemetry_honeycomb, "~> 0.5.0-rc.1"},
      {:hackney, "~> 1.17"},
      {:poison, "~> 4.0"}
    ]
  end
end
