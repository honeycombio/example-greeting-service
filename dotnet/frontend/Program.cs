using OpenTelemetry.Resources;
using OpenTelemetry.Trace;
using OpenTelemetry.Logs;
using OpenTelemetry.Exporter;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.

builder.Services.AddControllers();
builder.Services.AddHttpClient();
const string telemetrySourceName = "honeycomb.examples.frontend-service-dotnet";

var resourceBuilder = ResourceBuilder.CreateDefault()
    .AddService(builder.Configuration.GetValue<string>("Otlp:ServiceName"))
    .AddEnvironmentVariableDetector();

var configureOtlpExporter = (OtlpExporterOptions options) =>
{
    options.Endpoint = new Uri(builder.Configuration.GetValue<string>("Otlp:Endpoint"));
    var apiKey = builder.Configuration.GetValue<string>("Otlp:ApiKey");
    var dataset = builder.Configuration.GetValue<string>("Otlp:Dataset");
    options.Headers = $"x-honeycomb-team={apiKey},x-honeycomb-dataset={dataset}";
};

builder.Services.AddOpenTelemetry()
    .WithTracing(options => options
        .SetResourceBuilder(resourceBuilder)
        .AddSource(telemetrySourceName)
        .AddAspNetCoreInstrumentation()
        .AddHttpClientInstrumentation()
        .AddOtlpExporter(configureOtlpExporter));

builder.Logging.AddOpenTelemetry(options => options
    .SetResourceBuilder(resourceBuilder)
    .AddOtlpExporter(configureOtlpExporter)
);

// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();
builder.Services.AddSingleton(TracerProvider.Default.GetTracer(builder.Configuration.GetValue<string>("Otlp:ServiceName")));

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();

app.UseAuthorization();

app.MapControllers();

app.Run();
