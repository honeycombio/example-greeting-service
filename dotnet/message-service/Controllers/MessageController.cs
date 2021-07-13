using System;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
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

        private readonly Tracer _tracer;
        public MessageController(Tracer tracer)
        {
            _tracer = tracer;
        }

        [HttpGet]
        public async Task<string> GetAsync()
        {
            using (var span = _tracer.StartActiveSpan("Getting Message"))
            {
                span.SetAttribute("testAttribute", "Message");
                var message = await DetermineMessage();
                return message;
            }
        }

        private static async Task<string> DetermineMessage()
        {
            await SleepAwhile();
            var rng = new Random();
            var i = rng.Next(0, 9);
            return Messages[i];
        }

        private static async Task SleepAwhile()
        {
            await Task.Delay(50);
        }
    }
}
