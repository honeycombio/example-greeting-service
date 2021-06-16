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
        private readonly Tracer _tracer = TracerProvider.Default.GetTracer(Startup.ActivitySourceName);

        private static readonly int[] Years =
        {
            2015, 2016, 2017, 2018, 2019, 2020
        };

        [HttpGet]
        public async Task<int> GetAsync()
        {
            // using var activity = Startup.ActivitySource.StartActivity("DetermineYear");
            // activity?.SetTag("banana", 1);
            using var parentSpan = _tracer.StartActiveSpan("DetermineYear");
            parentSpan.SetAttribute("papaya", 1);
            var year = await DetermineYear();
            return year;
        }

        private async Task<int> DetermineYear()
        {
            await SleepAwhile();
            var rng = new Random();
            var i = rng.Next(0, 5);
            return Years[i];
        }

        private async Task SleepAwhile()
        {
            // using var activity = Startup.ActivitySource.StartActivity("Sleep");
            // activity?.SetTag("banana", 2);
            using var childSpan = _tracer.StartActiveSpan("Sleep");
            childSpan.SetAttribute("papaya", 2);
            await Task.Delay(50);
        }
    }
}
