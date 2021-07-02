FROM elixir:1.12
WORKDIR /app
COPY mix.* /app
COPY lib /app/lib
COPY config /app/config
RUN mix local.hex --force
RUN mix local.rebar --force
RUN mix deps.get
RUN mix deps.compile

EXPOSE 9000
CMD mix run --no-halt
