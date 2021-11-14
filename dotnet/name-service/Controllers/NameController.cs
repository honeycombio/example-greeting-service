using Microsoft.AspNetCore.Mvc;
using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;
using OpenTelemetry.Trace;
using OpenTelemetry;

namespace name_service.Controllers
{
    [Route("[controller]")]
    [ApiController]
    public class NameController : ControllerBase
    {
        private readonly IHttpClientFactory _clientFactory;
        private readonly Tracer _tracer;

        private static readonly Dictionary<int, string[]> _namesByYear = new()
        {
            { 2015, new[] { "sophia", "jackson", "emma", "aiden", "olivia", "liam", "ava", "lucas", "mia", "noah" } },
            { 2016, new[] { "sophia", "jackson", "emma", "aiden", "olivia", "lucas", "ava", "liam", "mia", "noah" } },
            { 2017, new[] { "sophia", "jackson", "olivia", "liam", "emma", "noah", "ava", "aiden", "isabella", "lucas" } },
            { 2018, new[] { "sophia", "jackson", "olivia", "liam", "emma", "noah", "ava", "aiden", "isabella", "caden" } },
            { 2019, new[] { "sophia", "liam", "olivia", "jackson", "emma", "noah", "ava", "aiden", "aria", "grayson" } },
            { 2020, new[] { "olivia", "noah", "emma", "liam", "ava", "elijah", "isabella", "oliver", "sophia", "lucas" } },
        };

        public NameController(IHttpClientFactory clientFactory)
        {
            _clientFactory = clientFactory;
            _tracer = TracerProvider.Default.GetTracer(Startup.TelemetrySourceName);
        }

        private static string GetYearEndpoint()
        {
            var yearEndpoint = Environment.GetEnvironmentVariable("YEAR_ENDPOINT");
            if (string.IsNullOrWhiteSpace(yearEndpoint))
            {
                return "http://localhost:6001/year";
            }
            else
            {
                return $"http://{yearEndpoint}/year";
            }
        }

        [HttpGet]
        public async Task<string> GetAsync()
        {
            var currentSpan = Tracer.CurrentSpan;
            currentSpan.SetAttribute("apple", 1);
            Baggage.SetBaggage("avocado", "12");

            var year = await GetYear();

            using var nameLookupSpan = _tracer.StartActiveSpan("📖 look up name based on year ✨");
            var name = "OH NO!";
            if (year != 0)
            {
                var rng = new Random();
                var i = rng.Next(0, 9);
                name = _namesByYear[year][i];
            }

            nameLookupSpan.SetAttribute("app.name", name);
            Baggage.SetBaggage("app.name", name);
            return name;
        }

        private async Task<int> GetYear()
        {
            using var yearServiceCallSpan = _tracer.StartActiveSpan("✨ call /year ✨");
            var request = new HttpRequestMessage(HttpMethod.Get, GetYearEndpoint());
            var client = _clientFactory.CreateClient();
            var response = await client.SendAsync(request);
            if (!response.IsSuccessStatusCode) return 0;
            var yearString = await response.Content.ReadAsStringAsync();
            return int.Parse(yearString);
        }
    }
}