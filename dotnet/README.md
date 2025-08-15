## .NET greeting service(s)

This is a .NET implementation of the example greeting service; 4 microservices that do some fancy greetin'.

You'll note that instrumentation is done differently depending on the project:

* `frontend` is a service built with .NET 6 that uses the standard OpenTelemetry libraries to configure instrumentation.
* `message-service` is a service built with .NET 5 that uses the standard OpenTelemetry libraries.
* `year-service` is a service built as a minimal .NET 6 API that uses the standard OpenTelemetry libraries to configure instrumentation.
* `name-service` is a service built with .NET 5 that uses the standard OpenTelemetry libraries.

The goal here is to demonstrate that you can configure instrumentation in several ways with .NET but still send data to Honeycomb.