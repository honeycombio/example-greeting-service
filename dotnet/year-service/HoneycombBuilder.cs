using System;
using OpenTelemetry;
using OpenTelemetry.Resources;
using OpenTelemetry.Trace;

namespace year_service
{
    public class HoneycombBuilder
    {
        private readonly string _serviceName;
        private readonly string _apiKey;
        private readonly string _dataset;
        private readonly string _endpoint;

        public HoneycombBuilder(string serviceName, string dataset, string apiKey, string endpoint)
        {
            this._serviceName = serviceName;
            this._dataset = dataset;
            this._apiKey = apiKey;
            _endpoint = endpoint;
        }

        public TracerProvider Build()
        {
            return Sdk.CreateTracerProviderBuilder()
                .SetResourceBuilder(ResourceBuilder.CreateDefault()
                    .AddService(_serviceName))
                .AddAspNetCoreInstrumentation()
                .AddHttpClientInstrumentation()
                .AddSource("honeycomb.examples.year-service-dotnet")
                .AddOtlpExporter(options =>
                {
                    options.Endpoint = new Uri(_endpoint);
                    options.Headers = $"x-honeycomb-team={_apiKey},x-honeycomb-dataset={_dataset}";
                })
                .Build();
        }

        public void Register(TracerProviderBuilder builder)
        {
            builder.SetResourceBuilder(ResourceBuilder.CreateDefault()
                    .AddService(_serviceName))
                .AddSource("honeycomb.examples.year-service-dotnet")
                .AddAspNetCoreInstrumentation()
                .AddHttpClientInstrumentation()
                .AddOtlpExporter(options =>
                {
                    options.Endpoint = new Uri(_endpoint);
                    options.Headers = $"x-honeycomb-team={_apiKey},x-honeycomb-dataset={_dataset}";
                });
        }
    }
}
