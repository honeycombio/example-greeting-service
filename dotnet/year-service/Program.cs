using OpenTelemetry.Resources;
using OpenTelemetry.Trace;

var builder = WebApplication.CreateBuilder(args);

const string telemetrySourceName = "honeycomb.examples.year-service-dotnet";

var app = builder.Build();
int[] years =
{
    2015, 2016, 2017, 2018, 2019, 2020
};

app.MapGet("/year", () =>
{
    var rng = new Random();
    var i = rng.Next(years.Length - 1);
    var year = years[i];
    return year;
});

app.Run();