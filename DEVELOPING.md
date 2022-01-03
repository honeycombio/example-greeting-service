# developing

## Environment Variables

There are a lot of environment variables.
A handy way to set these is to use [direnv](https://github.com/direnv/direnv)

Setup the hook for your shell by following the appropriate steps in [Setup](https://github.com/direnv/direnv/blob/master/docs/hook.md).

In the top-level directory of example-greeting-service, create a file called `.envrc` with the following environment variables:

```shell
export HONEYCOMB_API_KEY=
export HONEYCOMB_DATASET=greetings
export HONEYCOMB_API=https://api.honeycomb.io
export OTEL_EXPORTER_OTLP_ENDPOINT=${HONEYCOMB_API}
export OTEL_EXPORTER_OTLP_HEADERS="x-honeycomb-team=${HONEYCOMB_API_KEY},x-honeycomb-dataset=${HONEYCOMB_DATASET}"
```

**Tip**: Create and modify this and other `.envrc` files with vim; editing in VSCode seems to muddy the characters in unseen ways.

To use these environment variables, you'll need to `direnv allow .` whenever the file changes.
Consider adding an alias for your shell to more easiliy/quickly type a shortened version `da`:

`alias da="direnv allow ."`

In each language subdirectory (ruby, python, etc), add another `.envrc` file.
This file can inherit the environment variables from the higher-level directory with `source_up`.
Add the `source_up` and `layout` needed based on the language (**NOTE**: not recommended to add layout for node, `node_modules` should handle this for you):

```shell
# ruby
source_up
layout ruby
```

```shell
# python
source_up
layout python3
```

```shell
# node
source_up
```

## Using Tilt and Docker

Some languages have their own tiltfile and/or their own docker-compose.

### dotnet

- Use the tiltfile in the `dotnet` subdirectory
- Expects dotnet v6.0 or later (via Microsoft's [installation instructions](https://docs.microsoft.com/en-us/dotnet/core/install/))

[Microsoft's instructions](https://docs.microsoft.com/en-us/dotnet/core/install/linux-snap):

```bash
sudo snap install dotnet-sdk --classic --channel=6.0
sudo snap alias dotnet-sdk.dotnet dotnet
```

## Different dependencies

### go

For go services, use the go modules file to replace our beeline (and optionally the opentelemetry-sdk) with a copy on your local machine, so you can test in-progress changes:

```go
module main

go 1.14

replace github.com/honeycombio/beeline-go => /home/paul/projects/beeline-go

...
```

### ruby

In Ruby, modify the gemfile:

```ruby
source "https://rubygems.org"

gem "honeycomb-beeline", path: "/home/paul/projects/beeline-ruby"
```

### python

In Python, modify `pyproject.toml`:

```python
honeycomb-beeline = { path = "/Users/doug/src/github.com/honeycombio/beeline-python" }
# or
honeycomb-beeline = { path = "/Users/jamiedanielson/code/beeline-python/dist/honeycomb-beeline-2.17.42.tar.gz" }
```

### nodejs

In NodeJS, modify `package.json`:

```js
  "dependencies": {
    "honeycomb-beeline": "file:../../beeline-nodejs"
  }
```
