# Elixir Greeting Services

## Services
- Frontend: Service that calls the message, name and year services to create a greeting message served at `localhost:7777/greeting`.
- Message: Service that chooses a random greeting message.
- Name: Service that chooses a random name.
- Year: Service that returns a random year.

## How to run without Tilt/Docker

See top-level README.md for instructions on how to run services with Tilt and Docker (easiest).

### Install Elixir
If using option 1 or 3 below, make sure you have Elixir installed. The preferred way to install Elixir is with [asdf](https://asdf-vm.com/guide/getting-started.html#_1-install-dependencies).

Once you have asdf set up, run the following to install Elixir & Erlang. Check the version needed in the `.tool-versions` file in the root of this repo.

```shell
asdf plugin add erlang https://github.com/asdf-vm/asdf-erlang.git
```

```shell
asdf plugin-add elixir https://github.com/asdf-vm/asdf-elixir.git
```

```shell
asdf install erlang <insert-version-here>
asdf install elixir <insert-version-here>
```

### Run each service

Install dependencies in the service directory

```shell
mix local.hex --force && mix local.rebar --force && mix deps.get && mix deps.compile
```

In each service directory, run the mix command that starts the service. 

Frontend:

```shell
mix phx.server
```

Message/Name/Year:

```shell
mix run --no-halt
```

## See it in action

`curl localhost:7777/greeting` for greeting

`curl localhost:9000/message` for message only

`curl localhost:8000/name` for name only

`curl localhost:6001` for year only

## Troubleshooting

### Port in use error

If an error comes up about a port in use, either tilt or docker might still be running in the background.

To shut down docker services run `docker-compose down`

To shut down tilt services run `tilt down`
