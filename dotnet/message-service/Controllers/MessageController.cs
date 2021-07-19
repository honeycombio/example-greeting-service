using System;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using OpenTelemetry.Trace;

namespace message_service.Controllers
{
    using StackExchange.Redis;

    [Route("[controller]")]
    [ApiController]
    public class MessageController : ControllerBase
    {
        private static readonly string[] Messages =
        {
            "how are you?", "how are you doing?", "what's good?", "what's up?", "how do you do?", "sup?",
            "good day to you", "how are things?", "howzit?", "woohoo",
        };

        private readonly Tracer _tracer;
        private readonly IConnectionMultiplexer _redisConnection;

        public MessageController(Tracer tracer, IConnectionMultiplexer redisConnection)
        {
            _tracer = tracer;
            _redisConnection = redisConnection;
        }

        [HttpGet]
        public async Task<string> GetAsync()
        {
            using var span = _tracer.StartActiveSpan("Getting Message");
            {
                span.SetAttribute("testAttribute", "Message");
                var message = await DetermineMessage();
                return message;
            }
        }

        private async Task<string> DetermineMessage()
        {
            IDatabase db = _redisConnection.GetDatabase();
            var rng = new Random();
            var i = rng.Next(0, 9);
            db.StringSet("message", Messages[i]);

            await SleepAwhile();
            return db.StringGet("message");
        }

        private static async Task SleepAwhile()
        {
            await Task.Delay(50);
        }
    }
}