# greeting-service

[![OSS Lifecycle](https://img.shields.io/osslifecycle/honeycombio/example-greeting-service)](https://github.com/honeycombio/home/blob/main/honeycomb-oss-lifecycle-and-practices.md)

This is a ridiculously overengineered greeting service!
Its sole purpose is to create a greeting composed of a name and a message.
The name will be one of the 10 most popular baby names in the USA in a particular year.
The year will be chosen at random.
The message will be chosen from a pre-canned list of 10 or so messages.

The system is composed of several services, each with a specific responsibility:

- frontend - This service acts as an API aggregation layer.
  It depends on the name service and the message service.
- name-service - This service returns a random name. It depends on the year service.
- year-service - This service returns a random year.
- message-service - This service returns a random message from a list.

Each service can be written in any language, and they should use a combination of instrumentation techniques.
Some will be instrumented with a Honeycomb beeline and some with other standards, like OpenTelemetry.
The idea is to use these to test compatibility between various types of services.
The motivating use case is testing trace context header compatibility between Honeycomb and OpenTelemetry.

**.NET**
- [frontend](/dotnet/frontend) using [Honeycomb.OpenTelemetry](https://github.com/honeycombio/honeycomb-opentelemetry-dotnet)
- [message-service](/dotnet/message-service) using [Honeycomb.OpenTelemetry](https://github.com/honeycombio/honeycomb-opentelemetry-dotnet)
- [name-service](/dotnet/name-service) using [OpenTelemetry](https://github.com/open-telemetry/opentelemetry-dotnet)
- [year-service](/dotnet/year-service) using [OpenTelemetry](https://github.com/open-telemetry/opentelemetry-dotnet)

**Elixir**
- [frontend](/elixir/frontend) using [opentelemetry_phoenix](https://github.com/open-telemetry/opentelemetry-erlang-contrib/tree/main/instrumentation/opentelemetry_phoenix)
- [message](/elixir/message) using [opentelemetry_honeycomb](https://github.com/garthk/opentelemetry_honeycomb)
- [name](/elixir/name) using [opentelemetry_api](https://github.com/open-telemetry/opentelemetry-erlang)
- [year](/elixir/year) using [opentelemetry_api](https://github.com/open-telemetry/opentelemetry-erlang)

**Go**
- [frontend](/golang/frontend) using [opentelemetry-go](https://github.com/open-telemetry/opentelemetry-go)
- [message-service](/golang/message-service) using [beeline-go](https://github.com/honeycombio/beeline-go)
- [name-service](/golang/name-service) using [opentelemetry-go](https://github.com/open-telemetry/opentelemetry-go)
- [year-service](/golang/year-service) using [opentelemetry-go](https://github.com/open-telemetry/opentelemetry-go)

**Java**
- [frontend](/java/frontend) using [opentelemetry-javaagent](https://github.com/open-telemetry/opentelemetry-java-instrumentation)
- [message-service](/java/message-service) using [honeycomb-opentelemetry-javaagent](https://github.com/honeycombio/honeycomb-opentelemetry-java)
- [name-service](/java/name-service) using [honeycomb-opentelemetry-javaagent](https://github.com/honeycombio/honeycomb-opentelemetry-java)
- [year-service](/java/year-service) using [honeycomb-opentelemetry-skd](https://github.com/honeycombio/honeycomb-opentelemetry-java)

**Node**
- [frontend](/node/frontend) using [beeline-nodejs](https://github.com/honeycombio/beeline-nodejs)
- [message-service](/node/message-service) using [beeline-nodejs](https://github.com/honeycombio/beeline-nodejs)
- [name-service](/node/name-service) using [beeline-nodejs](https://github.com/honeycombio/beeline-nodejs)
- [year-service](/node/year-service) using [opentelemetry-js](https://github.com/open-telemetry/opentelemetry-js)
  
**Python**
- [frontend](/python/frontend) using [beeline-python](https://github.com/honeycombio/beeline-python)
- [message-service](/python/message-service) using [beeline-python](https://github.com/honeycombio/beeline-python)
- [name-service](/python/name-service) using [opentelemetry-python](https://github.com/open-telemetry/opentelemetry-python)
- [year-service](/python/year-service) using [beeline-python](https://github.com/honeycombio/beeline-python)

**Ruby**
- [frontend](/ruby/frontend) using [opentelemetry-ruby](https://github.com/open-telemetry/opentelemetry-ruby)
- [message-service](/ruby/message-service) using [beeline-ruby](https://github.com/honeycombio/beeline-ruby)
- [name-service](/ruby/name-service) using [beeline-ruby](https://github.com/honeycombio/beeline-ruby)
- [year-service](/ruby/year-service) using [beeline-ruby](https://github.com/honeycombio/beeline-ruby)

**Web**
- A web app instrumented using [opentelemetry-js](https://github.com/open-telemetry/opentelemetry-js)

## Caveats

Port and host names are hardcoded.

- Frontend: Port 7000
- Name: Port 8000
- Year: Port 6001
- Message: Port 9000

## Setup

Check out [DEVELOPING.md](DEVELOPING.md) for some additional tips and language-specific details.

Each service reads its configuration from the environment.
Specific environment variables:

- `HONEYCOMB_API_KEY` - Your honeycomb API key
- `OTEL_EXPORTER_OTLP_ENDPOINT=https://api.honeycomb.io`

Some services use vanilla OTEL:

- `OTEL_EXPORTER_OTLP_HEADERS='x-honeycomb-team=api-key'`

If configuring non-prod API endpoint:

- `OTEL_EXPORTER_OTLP_ENDPOINT=https://api.some.place`
- `HONEYCOMB_API_ENDPOINT=https://api.some.place`

If using Classic Honeycomb, you'll also need a dataset and must include in the OTEL headers:

- `HONEYCOMB_DATASET` - The name of the dataset you want to write to
- `OTEL_EXPORTER_OTLP_HEADERS='x-honeycomb-team=api-key,x-honeycomb-dataset=greetings'`

## Running

### Server apps
There is a `Tiltfile` to run these services on a local host using <https://tilt.dev/>.
After installing Tilt, running `tilt up` should spin up all of the services.

The default setup runs the go services.

To run services in another supported language, add the language name after the tilt command:

```shell
tilt up node
```

List of supported languages

- `go`
- `py`
- `rb`
- `java`
- `dotnet`
- `node`
- `elixir`

It's also possible to run a combination of services in different languages, for example the following command would run each specific service mentioned along with the required services (collector, redis, curl greeting)

```shell
tilt up frontend-node message-go name-py year-rb
```

To configure a common set of services that are specific to ongoing development, or to override the default option of running all services in go, add a file `tilt_config.json` and specify a group or set of services.
This file is ignored by git so it can be developer specific and allows running `tilt up` without having to specify further arguments.


Example `tilt_config.json` to override go as the default service

```json
{
  "to-run": ["node"]
}
```

Example `tilt_config.json` to override the default with multiple services

```json
{
  "to-run": ["frontend-node", "message-go", "name-py", "year-rb"]
}
```

Once running, `curl localhost:7000/greeting` to get a greeting and a trace!

ctrl+c to kill the session, and `tilt down` to spin down all services.

### Client apps

To run the browser app inside of `/web` run

```shell
tilt up web node 
```

This will start up the browser app as well as all node backend services. The browser app makes requests to `http://localhost:7000/greeting` so there has to be a set of backend services running. It could also be any one of our other supported languages (e.g. `py`, `go` etc.)
