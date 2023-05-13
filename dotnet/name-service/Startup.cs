using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Microsoft.OpenApi.Models;
using OpenTelemetry.Resources;
using OpenTelemetry.Trace;
using System;

namespace name_service
{
    public class Startup
    {
        public const string TelemetrySourceName = "honeycomb.examples.name-service-dotnet";

        public Startup(IConfiguration configuration)
        {
            Configuration = configuration;
        }

        public IConfiguration Configuration { get; }

        // This method gets called by the runtime. Use this method to add services to the container.
        public void ConfigureServices(IServiceCollection services)
        {
            services.AddControllers();
            services.AddSwaggerGen(c =>
            {
                c.SwaggerDoc("v1", new OpenApiInfo { Title = "name_service", Version = "v1" });
            });

            services.AddHttpClient();
            services.AddOpenTelemetry()
                .WithTracing((builder => builder
                    .SetResourceBuilder(ResourceBuilder.CreateDefault()
                        .AddService(this.Configuration.GetValue<string>("Otlp:ServiceName"))
                        .AddEnvironmentVariableDetector()
                    )
                    .AddSource(TelemetrySourceName)
                    .AddAspNetCoreInstrumentation()
                    .AddHttpClientInstrumentation()
                    .AddOtlpExporter(options =>
                    {
                        options.Endpoint = new Uri(Configuration.GetValue<string>("Otlp:Endpoint"));
                        var apiKey = Configuration.GetValue<string>("Otlp:ApiKey");
                        var dataset = Configuration.GetValue<string>("Otlp:Dataset");
                        options.Headers = $"x-honeycomb-team={apiKey},x-honeycomb-dataset={dataset}";
                    })));
        }

        // This method gets called by the runtime. Use this method to configure the HTTP request pipeline.
        public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
        {
            if (env.IsDevelopment())
            {
                app.UseDeveloperExceptionPage();
                app.UseSwagger();
                app.UseSwaggerUI(c => c.SwaggerEndpoint("/swagger/v1/swagger.json", "name_service v1"));
            }

            app.UseRouting();

            app.UseEndpoints(endpoints => endpoints.MapControllers());
        }
    }
}
