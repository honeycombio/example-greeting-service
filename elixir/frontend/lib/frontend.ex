defmodule Frontend do
  def name do
    endpoint = System.get_env("NAME_ENDPOINT", "http://localhost:8000")
    HTTPoison.get!("#{endpoint}/year").body
  end

  def message do
    endpoint = System.get_env("MESSAGE_ENDPOINT", "http://localhost:9000")
    HTTPoison.get!("#{endpoint}/message").body
  end
end
