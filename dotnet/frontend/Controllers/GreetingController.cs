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

        [HttpGet]
        public async Task<string> GetAsync()
        {
            using (var span = _tracer.StartActiveSpan("Preparing Greeting"))
            {

                span.SetAttribute("testAttribute", "Greeting");
                Baggage.Current.SetBaggage("testBaggage", "Greetings");
                var httpClient = _clientFactory.CreateClient();
                var name = await GetNameAsync(httpClient);
                var message = await GetMessage(httpClient);

                return $"Hello {name}, {message}";
            }
        }

        private static async Task<string> GetNameAsync(HttpClient client)
        {
            var request = new HttpRequestMessage(HttpMethod.Get, "http://localhost:8000/name");
            var response = await client.SendAsync(request);
            if (!response.IsSuccessStatusCode) return "OH NO!";
            return await response.Content.ReadAsStringAsync();
        }

        private static async Task<string> GetMessage(HttpClient client)
        {
            var request = new HttpRequestMessage(HttpMethod.Get, "http://localhost:9000/message");
            var response = await client.SendAsync(request);
            if (!response.IsSuccessStatusCode) return "OH NO!";
            return await response.Content.ReadAsStringAsync();
        }
    }
}
