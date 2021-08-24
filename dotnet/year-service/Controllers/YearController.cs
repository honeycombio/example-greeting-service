using System;
using Microsoft.AspNetCore.Mvc;

namespace year_service.Controllers
{
    [Route("[controller]")]
    [ApiController]
    public class YearController : ControllerBase
    {
        private static readonly int[] _years =
        {
            2015, 2016, 2017, 2018, 2019, 2020
        };

        [HttpGet]
        public int GetYear()
        {
            using var activity = Startup.ActivitySource.StartActivity("🗓 get-a-year ✨");
            activity?.SetTag("banana", 1);
            var year = DetermineYear();
            return year;
        }

        private static int DetermineYear()
        {
            var rng = new Random();
            var i = rng.Next(_years.Length-1);
            return _years[i];
        }
    }
}
