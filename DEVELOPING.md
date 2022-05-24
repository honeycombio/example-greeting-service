# developing

## Environment Variables

There are a lot of environment variables.
A handy way to set these is to use [direnv](https://github.com/direnv/direnv)

Setup the hook for your shell by following the appropriate steps in [Setup](https://github.com/direnv/direnv/blob/master/docs/hook.md).

In the top-level directory of example-greeting-service, create a file called `.envrc` with the following environment variables:

```shell
export HONEYCOMB_API_KEY=
export HONEYCOMB_API=https://api.honeycomb.io
export HONEYCOMB_API_ENDPOINT=${HONEYCOMB_API}
export OTEL_EXPORTER_OTLP_ENDPOINT=${HONEYCOMB_API}
export OTEL_EXPORTER_OTLP_HEADERS="x-honeycomb-team=${HONEYCOMB_API_KEY}"
```

`SERVICE_NAME` is set in the Tiltfile when running `tilt up`.
If not using Tilt, the service name must be set.

If using Classic Honeycomb, a dataset must also be included as a variable and as part of the headers:

```shell
export HONEYCOMB_API_KEY=
export HONEYCOMB_DATASET=greetings
export HONEYCOMB_API=https://api.honeycomb.io
export HONEYCOMB_API_ENDPOINT=${HONEYCOMB_API}
export OTEL_EXPORTER_OTLP_ENDPOINT=${HONEYCOMB_API}
export OTEL_EXPORTER_OTLP_HEADERS="x-honeycomb-team=${HONEYCOMB_API_KEY},x-honeycomb-dataset=${HONEYCOMB_DATASET}"
```

**Tip**: Monterey Control Center uses port 7000 for AirPlay and this will break your ability to run EGS; you can turn it off in the System Settings at the bottom of the Sharing settings.

To use these environment variables, you'll need to `direnv allow .` whenever the file changes.
Consider adding an alias for your shell to more easily/quickly type a shortened version `da`:

`alias da="direnv allow ."`

In each language subdirectory (ruby, python, etc), add another `.envrc` file.
This file can inherit the environment variables from the higher-level directory with `source_up`.
Add the `source_up` and `layout` needed based on the language (**NOTE**: not recommended to add layout for node, `node_modules` should handle this for you):

```shell
# go
source_up
layout go
```

```shell
# java
source_up
layout java
```

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

### node

Node specific setup instructions can be found in the [README](node/README.md) in the `node` folder.

## Different dependencies

### go

For go services, use the go modules file to replace our beeline (and optionally the opentelemetry-sdk) with a copy on your local machine, so you can test in-progress changes:

```txt
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

```toml
honeycomb-beeline = { path = "/Users/doug/src/github.com/honeycombio/beeline-python", develop = true} }
# or
honeycomb-beeline = { path = "/Users/jamie/code/beeline-python/dist/honeycomb-beeline-2.17.42.tar.gz" }
```

### nodejs

In NodeJS, modify `package.json`:

```js
  "dependencies": {
    "honeycomb-beeline": "file:../../beeline-nodejs"
  }
  // or
  "dependencies": {
    "honeycomb-beeline": "file:/Users/jamie/code/beeline-nodejs"
  }
```
