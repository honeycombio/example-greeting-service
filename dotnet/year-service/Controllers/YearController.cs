using System;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using OpenTelemetry.Trace;

namespace year_service.Controllers
{
    [Route("[controller]")]
    [ApiController]
    public class YearController : ControllerBase
    {
        private static readonly int[] Years =
        {
            2015, 2016, 2017, 2018, 2019, 2020
        };

        private readonly Tracer _tracer;
        public YearController(Tracer tracer)
        {
            _tracer = tracer;
        }

        [HttpGet]
        public async Task<int> GetAsync()
        {
            using (var span = _tracer.StartActiveSpan("Determine Year"))
            {

                span.SetAttribute("testAttribute", "Year");
                var year = await DetermineYear();
                return year;
            }
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
            // NOTE: tracer doesn't work on static, Activity was removed
            // using var activity = Startup.ActivitySource.StartActivity("Sleep");
            // activity?.SetTag("banana", 2);
            await Task.Delay(50);
        }
    }
}
