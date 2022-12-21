## .NET greeting service(s)

This is a .NET implementation of the example greeting service; 4 microservices that do some fancy greetin'.

You'll note that instrumentation is done differently depending on the project:

* `frontend` is a service built with .NET 6 that uses [Honeycomb.OpenTelemetry](https://www.nuget.org/packages/Honeycomb.OpenTelemetry), our distribution of the .NET OpenTelemetry SDK. Configuration is simpler when you use this package.
* `message-service` is a service built with .NET 5 that uses the Honeycomb.OpenTelemetry distribution.
* `year-service` is a service built as a minimal .NET 6 API that uses the standard OpenTelemetry libraries to configure instrumentation. It's more lines of code.
* `name-service` is a service built with .NET 5 that uses the standard OpenTelemetry libraries.

The goal here is to demonstrate that you can configure instrumentation in several ways with .NET but still send data to Honeycomb.

### Run the services

To run the services, you can either run `tilt up dotnet` from the root directory or run `docker compose up` from this directory. It is not recommended to try and run the services outside of docker since they (deliberately) run different versions of .NET.
