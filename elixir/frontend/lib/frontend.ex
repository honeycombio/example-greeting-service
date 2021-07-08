defmodule Frontend do
  require OpenTelemetry.Tracer, as: Tracer

  def name do
    Tracer.with_span "✨ call /name ✨" do
      endpoint = System.get_env("NAME_ENDPOINT", "http://localhost:8000")
      headers = :otel_propagator.text_map_inject([])
      HTTPoison.get!("#{endpoint}/name", headers).body
    end
  end

  def message do
    Tracer.with_span "✨ call /message ✨" do
      endpoint = System.get_env("MESSAGE_ENDPOINT", "http://localhost:9000")
      headers = :otel_propagator.text_map_inject([])
      HTTPoison.get!("#{endpoint}/message", headers).body
    end
  end
end
