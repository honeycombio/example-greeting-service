using System.Diagnostics;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Microsoft.OpenApi.Models;
using Honeycomb.OpenTelemetry;

namespace year_service
{
    public class Startup
    {
        private const string ActivitySourceName = "honeycomb.examples.year-service-dotnet";
        public static readonly ActivitySource ActivitySource = new(ActivitySourceName);

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
                c.SwaggerDoc("v1", new OpenApiInfo {Title = "year_service", Version = "v1"});
            });

            // services.AddOpenTelemetryTracing(builder => builder
            //     .SetResourceBuilder(ResourceBuilder.CreateDefault()
            //         .AddService(this.Configuration.GetValue<string>("Otlp:ServiceName")))
            //     .AddSource(ActivitySourceName)
            //     .AddAspNetCoreInstrumentation()
            //     .AddHttpClientInstrumentation()
            //     .AddOtlpExporter(options =>
            //     {
            //         options.Endpoint = new Uri(Configuration.GetValue<string>("Otlp:Endpoint"));
            //         var apiKey = Configuration.GetValue<string>("Otlp:ApiKey");
            //         var dataset = Configuration.GetValue<string>("Otlp:Dataset");
            //         options.Headers = $"x-honeycomb-team={apiKey},x-honeycomb-dataset={dataset}";
            //     }));

            // services.AddHoneycombOpenTemeletry(); // defualt to using env vars
            services.AddHoneycombOpenTemeletry(builder => // provide options in-line
            {
                builder.WithServiceName(Configuration.GetValue<string>("Otlp:ServiceName"));
                builder.WithAPIKey(Configuration.GetValue<string>("Otlp:ApiKey"));
                builder.WithDataset(Configuration.GetValue<string>("Otlp:Dataset"));
                builder.WithEndpoint(Configuration.GetValue<string>("Otlp:Endpoint"));
                builder.WithSources(new string[] {ActivitySourceName});

                // TODO: auto bind config values to automate ^^
                // TODO: hide .Build(), we should only use in the sdk, not here
            });
        }

        // This method gets called by the runtime. Use this method to configure the HTTP request pipeline.
        public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
        {
            if (env.IsDevelopment())
            {
                app.UseDeveloperExceptionPage();
                app.UseSwagger();
                app.UseSwaggerUI(c => c.SwaggerEndpoint("/swagger/v1/swagger.json", "year_service v1"));
            }

            app.UseRouting();

            app.UseEndpoints(endpoints => { endpoints.MapControllers(); });
        }
    }
}
