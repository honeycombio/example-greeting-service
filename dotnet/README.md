## .NET greeting service(s)

This is a .NET implementation of the example greeting service; 4 microservices that do some fancy greetin'.

You'll note that instrumentation is done differently depending on the project:

* `frontend` and `message-service` use [Honeycomb.OpenTelemetry](https://www.nuget.org/packages/Honeycomb.OpenTelemetry), our distribution of the .NET OpenTelemetry SDK. Configuration is simpler when you use this package.
* `year-service` uses the standard .NET OpenTelemetry APIs to configure instrumentation. It's more lines of code.
* `name-service` uses the standard .NET OpenTelemetry APIs but also configures the ASP.NET Core instrumentation a little more, too. It's even more lines of code.

The goal here is to demonstrate that you can configure instrumentation in several ways with .NET but still send data to Honeycomb.