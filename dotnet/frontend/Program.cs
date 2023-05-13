using OpenTelemetry.Trace;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.

builder.Services.AddControllers();
builder.Services.AddHttpClient();
var honeycombOptions = builder.Configuration.GetHoneycombOptions();
builder.Services.AddOpenTelemetry().WithTracing(otelBuilder =>
{
    otelBuilder
        .AddHoneycomb(honeycombOptions)
        .AddCommonInstrumentations()
        .AddAspNetCoreInstrumentationWithBaggage();
});
builder.Services.AddSingleton(TracerProvider.Default.GetTracer(honeycombOptions.ServiceName));

// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

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
