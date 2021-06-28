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
        private readonly Tracer tracer;

        public YearController(Tracer tracer)
        {
            this.tracer = tracer;
        }

        [HttpGet]
        public async Task<int> GetAsync()
        {
            using (var span = tracer.StartActiveSpan("DetermineYear"))
            {
                span.SetAttribute("banana", 1);
                using (var delaySpan = tracer.StartActiveSpan("Sleep"))
                {
                    delaySpan.SetAttribute("banana", 2);
                    await Task.Delay(50);
                }
                var rng = new Random();
                return Years[rng.Next(0, 5)];
            }
        }
    }
}
