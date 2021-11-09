using OpenTelemetry.Resources;
using OpenTelemetry.Trace;
using System.Diagnostics;

var builder = WebApplication.CreateBuilder(args);

const string activitySourceName = "honeycomb.examples.year-service-dotnet";
ActivitySource activitySource = new(activitySourceName);

builder.Services.AddOpenTelemetryTracing(providerBuilder => providerBuilder
    .SetResourceBuilder(ResourceBuilder.CreateDefault()
        .AddService(builder.Configuration.GetValue<string>("Otlp:ServiceName"))
        .AddEnvironmentVariableDetector()
    )
    .AddSource(activitySourceName)
    .AddAspNetCoreInstrumentation()
    .AddHttpClientInstrumentation()
    .AddOtlpExporter(options =>
    {
        options.Endpoint = new Uri(builder.Configuration.GetValue<string>("Otlp:Endpoint"));
        var apiKey = builder.Configuration.GetValue<string>("Otlp:ApiKey");
        var dataset = builder.Configuration.GetValue<string>("Otlp:Dataset");
        options.Headers = $"x-honeycomb-team={apiKey},x-honeycomb-dataset={dataset}";
    }));

var app = builder.Build();
int[] years =
{
    2015, 2016, 2017, 2018, 2019, 2020
};

app.MapGet("/year", () =>
{
    using var activity = activitySource.StartActivity("🗓 get-a-year ✨");
    activity?.SetTag("banana", 1);
    var rng = new Random();
    var i = rng.Next(years.Length - 1);
    var year = years[i];
    return year;
});

app.Run();