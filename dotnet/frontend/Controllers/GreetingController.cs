using Microsoft.AspNetCore.Mvc;
using OpenTelemetry;
using OpenTelemetry.Trace;
using System;
using System.Net.Http;
using System.Threading.Tasks;

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
            if (string.IsNullOrWhiteSpace(nameEndpoint))
            {
                return "http://localhost:8000/name";
            }
            else
            {
                return $"http://{nameEndpoint}/name";
            }
        }

        private static string GetMessageEndpoint()
        {
            var messageEndpoint = Environment.GetEnvironmentVariable("MESSAGE_ENDPOINT");
            if (string.IsNullOrWhiteSpace(messageEndpoint))
            {
                return "http://localhost:9000/message";
            }
            else
            {
                return $"http://{messageEndpoint}/message";
            }
        }

        [HttpGet]
        public async Task<string> GetAsync()
        {
            Console.WriteLine("Starting GetMessage!");
            var httpClient = _clientFactory.CreateClient();
            var name = await GetNameAsync(httpClient);
            var message = await GetMessageAsync(httpClient);

            if(name.Length == 0) {
                Console.WriteLine("Zero length name should never happen");
            }

            var rv = $"Hello {name}, {message}";
            Console.WriteLine("Finishing GetMessage " + rv );
            return rv;
        }

        private async Task<string> GetNameAsync(HttpClient client)
        {
            var request = new HttpRequestMessage(HttpMethod.Get, GetNameEndpoint());
            var response = await client.SendAsync(request);
            if (!response.IsSuccessStatusCode)
            {
                return "OH NO!";
            }

            return await response.Content.ReadAsStringAsync();
        }

        private async Task<string> GetMessageAsync(HttpClient client)
        {
            var request = new HttpRequestMessage(HttpMethod.Get, GetMessageEndpoint());
            var response = await client.SendAsync(request);
            if (!response.IsSuccessStatusCode)
            {
                return "OH NO!";
            }

            return await response.Content.ReadAsStringAsync();
        }
    }
}