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
            var request = new HttpRequestMessage(HttpMethod.Get, "http://localhost:5001/name");
            var response = await client.SendAsync(request);
            if (!response.IsSuccessStatusCode) return "OH NO!";
            return await response.Content.ReadAsStringAsync();
        }

        private static async Task<string> GetMessage(HttpClient client)
        {
            var request = new HttpRequestMessage(HttpMethod.Get, "http://localhost:5002/message");
            var response = await client.SendAsync(request);
            if (!response.IsSuccessStatusCode) return "OH NO!";
            return await response.Content.ReadAsStringAsync();
        }
    }
}
