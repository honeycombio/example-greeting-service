using System;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using OpenTelemetry.Trace;

namespace message_service.Controllers
{
    [Route("[controller]")]
    [ApiController]
    public class MessageController : ControllerBase
    {
        private static readonly string[] Messages =
        {
            "how are you?", "how are you doing?", "what's good?", "what's up?", "how do you do?",
            "sup?", "good day to you", "how are things?", "howzit?", "woohoo",
        };

        // private readonly ILogger<MessageController> _logger;
        // private readonly Tracer _tracer;

        // public MessageController(ILogger<MessageController> logger, Tracer tracer)
        // {
        //     _logger = logger;
        //     _tracer = tracer;
        // }

        [HttpGet]
        public async Task<string> GetAsync()
        {
            var message = await DetermineMessage();
            return message;
        }

        private async Task<string> DetermineMessage()
        {
            // using (var messageSpan = _tracer.StartActiveSpan("Determine Message"))
            {
                await SleepAwhile();
                var rng = new Random();
                var i = rng.Next(0, 9);
                return Messages[i];
            }
        }

        private async Task SleepAwhile()
        {
            // using (var sleepSpan = _tracer.StartActiveSpan("sleep"))
            {
                // sleepSpan.SetAttribute("delay_ms", 100);
                await Task.Delay(50);
            }
        }
    }
}
