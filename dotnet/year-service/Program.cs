using OpenTelemetry.Resources;
using OpenTelemetry.Trace;

var builder = WebApplication.CreateBuilder(args);

const string telemetrySourceName = "honeycomb.examples.year-service-dotnet";

builder.Services.AddOpenTelemetryTracing(providerBuilder => providerBuilder
    .SetResourceBuilder(ResourceBuilder.CreateDefault()
        .AddService(builder.Configuration.GetValue<string>("Otlp:ServiceName"))
        .AddEnvironmentVariableDetector()
    )
    .AddSource(telemetrySourceName)
    .AddAspNetCoreInstrumentation()
    .AddHttpClientInstrumentation()
    .AddOtlpExporter(options =>
    {
        options.Endpoint = new Uri(builder.Configuration.GetValue<string>("Otlp:Endpoint"));
        var apiKey = builder.Configuration.GetValue<string>("Otlp:ApiKey");
        var dataset = builder.Configuration.GetValue<string>("Otlp:Dataset");
        options.Headers = $"x-honeycomb-team={apiKey},x-honeycomb-dataset={dataset}";
    }));

Tracer tracer = TracerProvider.Default.GetTracer(telemetrySourceName);
var app = builder.Build();
int[] years =
{
    2015, 2016, 2017, 2018, 2019, 2020
};

app.MapGet("/year", () =>
{
    using var span = tracer.StartActiveSpan("🗓 get-a-year ✨");
    span.SetAttribute("banana", 1);
    var rng = new Random();
    var i = rng.Next(years.Length - 1);
    var year = years[i];
    return year;
});

app.Run();