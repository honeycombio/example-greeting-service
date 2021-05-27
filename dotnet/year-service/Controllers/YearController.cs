using System;
using System.Diagnostics;
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
            using (Activity activity = Startup.ActivitySource.StartActivity("DetermineYear"))
            {
                activity?.SetTag("banana", 1);
                var year = await DetermineYear();
                return year;
            }
        }

        private static async Task<int> DetermineYear()
        {
            await Task.Delay(50);
            var rng = new Random();
            var i = rng.Next(0, 5);
            return Years[i];
        }
    }
}
