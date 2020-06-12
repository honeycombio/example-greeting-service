# greeting-service

This is a ridiculously overengineered greeting service! It's sole purpose is to
create a greeting, composed of a name and a message. The name will be one of
the 10 most popular baby names in the USA in a particular year. The year will
be chosen at random. The message will be chosen from a pre-canned list of 10 or
so messages.

The system is composed of several services, each with a specific responsibility:

* frontend - This service acts as an API aggregation layer. It depends on the
             name service and the message service.
* name-service - This service returns a random name. It depends on the year service.
* year-service - This service returns a random year.
* message-service - This service returns a random message from a list.

Each service can be written in any language, and they should use a combination
of instrumentation techniques. Some will be instrumented with a Honeycomb
beeline and some with other standards, like OpenTelemetry. The idea is to use
these to test compabitility between various types of services. The motivating
use case is testing trace context header compatibility between Honeycomb and
OpenTelemetry.

Each service reads it's configuration from the environment. Specific environment
variables:

* HONEYCOMB_WRITE_KEY - Your honeycomb API key
* HONEYCOMB_DATASET - The name of the dataset you want to write to

### Caveats

Port and host names hardcoded. We should use docker-compose or something for
this. These services are stupid! 

All of these statements are true!

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
