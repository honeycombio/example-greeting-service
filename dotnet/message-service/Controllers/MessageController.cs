using Microsoft.AspNetCore.Mvc;
using OpenTelemetry.Trace;
using System.Threading.Tasks;
using System;


namespace message_service.Controllers
{
    using StackExchange.Redis;

    [Route("[controller]")]
    [ApiController]
    public class MessageController : ControllerBase
    {
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
            var message = await DetermineMessage();
            return message;
        }

        private async Task<string> DetermineMessage()
        {
            Console.WriteLine("DetermineMessage");
            var db = _redisConnection.GetDatabase();
            var message = "generic hello";
            var result = await db.SetRandomMemberAsync("messages");
            if (result.IsNull)
            {
                Console.WriteLine("Message was empty");
            }
            else
            {
                message = result.ToString();
            }
            Console.WriteLine("DetermineMessage returned " + message);
            return message;
        }
    }
}