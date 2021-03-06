using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Net.Http;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;

namespace name_service.Controllers
{
    [Route("[controller]")]
    [ApiController]
    public class NameController : ControllerBase
    {
        private readonly IHttpClientFactory _clientFactory;

        public NameController(IHttpClientFactory clientFactory)
        {
            _clientFactory = clientFactory;
        }

        private static readonly Dictionary<int, string[]> NamesByYear = new()
        {
            {2015, new[] {"sophia", "jackson", "emma", "aiden", "olivia", "liam", "ava", "lucas", "mia", "noah"}},
            {2016, new[] {"sophia", "jackson", "emma", "aiden", "olivia", "lucas", "ava", "liam", "mia", "noah"}},
            {2017, new[] {"sophia", "jackson", "olivia", "liam", "emma", "noah", "ava", "aiden", "isabella", "lucas"}},
            {2018, new[] {"sophia", "jackson", "olivia", "liam", "emma", "noah", "ava", "aiden", "isabella", "caden"}},
            {2019, new[] {"sophia", "liam", "olivia", "jackson", "emma", "noah", "ava", "aiden", "aria", "grayson"}},
            {2020, new[] {"olivia", "noah", "emma", "liam", "ava", "elijah", "isabella", "oliver", "sophia", "lucas"}},
        };


        [HttpGet]
        public async Task<string> GetAsync()
        {
            var current = Activity.Current;
            current?.AddTag("apple", 1);
            current?.AddBaggage("avocado", "12");

            var request = new HttpRequestMessage(HttpMethod.Get, "http://localhost:6001/year");
            var client = _clientFactory.CreateClient();
            var response = await client.SendAsync(request);
            if (!response.IsSuccessStatusCode) return "OH NO!";
            var yearString = await response.Content.ReadAsStringAsync();
            var year = int.Parse(yearString);

            var rng = new Random();
            var i = rng.Next(0, 9);
            return NamesByYear[year][i];
        }
    }
}
