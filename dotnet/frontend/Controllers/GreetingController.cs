using System;
using System.Diagnostics;
using System.Net.Http;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;

namespace frontend.Controllers
{
    [Route("[controller]")]
    [ApiController]
    public class GreetingController : ControllerBase
    {
        private readonly IHttpClientFactory _clientFactory;

        public GreetingController(IHttpClientFactory clientFactory)
        {
            _clientFactory = clientFactory;
        }

        private static string GetNameEndpoint()
        {
            var nameEndpoint = Environment.GetEnvironmentVariable("NAME_ENDPOINT");
            if (nameEndpoint == null)
            {
                return "http://localhost:8000/name";
            } else {
                return nameEndpoint + "/name";
            }
        }

        private static string GetMessageEndpoint()
        {
            var messageEndpoint = Environment.GetEnvironmentVariable("MESSAGE_ENDPOINT");
            if (messageEndpoint == null)
            {
                return "http://localhost:9000/message";
            } else {
                return messageEndpoint + "/message";
            }
        }

        [HttpGet]
        public async Task<string> GetAsync()
        {
            var current = Activity.Current;
            current?.AddTag("apple", 1);
            current?.AddBaggage("avocado", "12");

            var httpClient = _clientFactory.CreateClient();
            var name = await GetNameAsync(httpClient);
            var message = await GetMessage(httpClient);

            return $"Hello {name}, {message}";
        }

        private static async Task<string> GetNameAsync(HttpClient client)
        {
            var request = new HttpRequestMessage(HttpMethod.Get, GetNameEndpoint());
            var response = await client.SendAsync(request);
            if (!response.IsSuccessStatusCode) return "OH NO!";
            return await response.Content.ReadAsStringAsync();
        }

        private static async Task<string> GetMessage(HttpClient client)
        {
            var request = new HttpRequestMessage(HttpMethod.Get, GetMessageEndpoint());
            var response = await client.SendAsync(request);
            if (!response.IsSuccessStatusCode) return "OH NO!";
            return await response.Content.ReadAsStringAsync();
        }
    }
}
