using System;
using System.Diagnostics;
using System.Net.Http;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using OpenTelemetry.Trace;
using OpenTelemetry;

namespace frontend.Controllers
{
    [Route("[controller]")]
    [ApiController]
    public class GreetingController : ControllerBase
    {
        private readonly IHttpClientFactory _clientFactory;
        private readonly Tracer _tracer;

        public GreetingController(IHttpClientFactory clientFactory, Tracer tracer)
        {
            _clientFactory = clientFactory;
            _tracer = tracer;
        }

        private static string GetNameEndpoint()
        {
            var nameEndpoint = Environment.GetEnvironmentVariable("NAME_ENDPOINT");
            if (nameEndpoint == null)
            {
                return "http://localhost:8000/name";
            } else {
                return "http://" + nameEndpoint + "/name";
            }
        }

        private static string GetMessageEndpoint()
        {
            var messageEndpoint = Environment.GetEnvironmentVariable("MESSAGE_ENDPOINT");
            if (messageEndpoint == null)
            {
                return "http://localhost:9000/message";
            } else {
                return "http://" + messageEndpoint + "/message";
            }
        }

        [HttpGet]
        public async Task<string> GetAsync()
        {
            using var span = _tracer.StartActiveSpan("Preparing Greeting");
            {
                span.SetAttribute("testAttribute", "Greeting");
                Baggage.Current.SetBaggage("testBaggage", "Greetings");
                var httpClient = _clientFactory.CreateClient();
                var name = await GetNameAsync(httpClient);
                var message = await GetMessageAsync(httpClient);

                using var render_span = _tracer.StartActiveSpan("🎨 render greeting ✨");
                {
                    return $"Hello {name}, {message}";
                }
            }
        }

        private async Task<string> GetNameAsync(HttpClient client)
        {
            using var span = _tracer.StartActiveSpan("✨ call /name ✨");
            var request = new HttpRequestMessage(HttpMethod.Get, GetNameEndpoint());
            var response = await client.SendAsync(request);
            if (!response.IsSuccessStatusCode) return "OH NO!";
            return await response.Content.ReadAsStringAsync();
        }

        private async Task<string> GetMessageAsync(HttpClient client)
        {
            using var span = _tracer.StartActiveSpan("✨ call /message ✨");
            var request = new HttpRequestMessage(HttpMethod.Get, GetMessageEndpoint());
            var response = await client.SendAsync(request);
            if (!response.IsSuccessStatusCode) return "OH NO!";
            return await response.Content.ReadAsStringAsync();
        }
    }
}
