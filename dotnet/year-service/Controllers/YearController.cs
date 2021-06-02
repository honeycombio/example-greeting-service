using System;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;

namespace year_service.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class YearController : ControllerBase
    {
        private static readonly int[] Years =
        {
            2015, 2016, 2017, 2018, 2019, 2020
        };

        [HttpGet]
        public async Task<int> GetAsync()
        {
            using var activity = Startup.ActivitySource.StartActivity("DetermineYear");
            activity?.SetTag("banana", 1);
            var year = await DetermineYear();
            return year;
        }

        private static async Task<int> DetermineYear()
        {
            await SleepAwhile();
            var rng = new Random();
            var i = rng.Next(0, 5);
            return Years[i];
        }

        private static async Task SleepAwhile()
        {
            using var activity = Startup.ActivitySource.StartActivity("Sleep");
            activity?.SetTag("banana", 2);
            await Task.Delay(50);
        }
    }
}
