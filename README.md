# greeting-service

[![OSS Lifecycle](https://img.shields.io/osslifecycle/honeycombio/example-greeting-service)](https://github.com/honeycombio/home/blob/main/honeycomb-oss-lifecycle-and-practices.md)

This is a ridiculously overengineered greeting service! It's sole purpose is to
create a greeting, composed of a name and a message. The name will be one of
the 10 most popular baby names in the USA in a particular year. The year will
be chosen at random. The message will be chosen from a pre-canned list of 10 or
so messages.

The system is composed of several services, each with a specific responsibility:

- frontend - This service acts as an API aggregation layer. It depends on the
  name service and the message service.
- name-service - This service returns a random name. It depends on the year service.
- year-service - This service returns a random year.
- message-service - This service returns a random message from a list.

Each service can be written in any language, and they should use a combination
of instrumentation techniques. Some will be instrumented with a Honeycomb
beeline and some with other standards, like OpenTelemetry. The idea is to use
these to test compabitility between various types of services. The motivating
use case is testing trace context header compatibility between Honeycomb and
OpenTelemetry.

Each service reads it's configuration from the environment. Specific environment
variables:

- HONEYCOMB_API_KEY - Your honeycomb API key
- HONEYCOMB_DATASET - The name of the dataset you want to write to
- OTEL_EXPORTER_OTLP_ENDPOINT=https://api.honeycomb.io

Some services use vanilla OTEL:
- OTEL_EXPORTER_OTLP_HEADERS='x-honeycomb-team=api-key,x-honeycomb-dataset=greetings'

If using dogfood:
- OTEL_EXPORTER_OTLP_ENDPOINT=https://api-dogfood.honeycomb.io
- HONEYCOMB_API_ENDPOINT=https://api-dogfood.honeycomb.io

### Caveats

Port and host names hardcoded. We should use docker-compose or something for
this. These services are stupid!

All of these statements are true!

### Running

There is a `Tiltfile` to run these services on a local host using https://tilt.dev/ - after installing Tilt, running `tilt up` should spin up all of the services.

#### Language-Specific notes

##### dotnet

- Use the tiltfile in the `dotnet` subdirectory
- Expects dotnet v6.0 or later (via Microsoft's [installation instructions](https://docs.microsoft.com/en-us/dotnet/core/install/))

[Microsoft's instructions](https://docs.microsoft.com/en-us/dotnet/core/install/linux-snap):

  ```bash
  sudo snap install dotnet-sdk --classic --channel=6.0
  sudo snap alias dotnet-sdk.dotnet dotnet
  ```

### Tips

For go services, use the go modules file to replace our beeline (and optionally the opentelemetry-sdk) with a copy on your local machine, so you can test in-progress changes:

```
module main

go 1.14

replace github.com/honeycombio/beeline-go => /home/paul/projects/beeline-go

...
```

Or in Ruby, modify the gemfile:

```
source "https://rubygems.org"

gem "honeycomb-beeline", path: "/home/paul/projects/beeline-ruby"
```

In Python, modify pyproject.toml:

```
honeycomb-beeline = { path = "/Users/doug/src/github.com/honeycombio/beeline-python" }
```
