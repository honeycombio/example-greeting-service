using Microsoft.AspNetCore.Mvc;
using OpenTelemetry.Trace;
using System.Threading.Tasks;

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
//            var currentSpan = Tracer.CurrentSpan;
//            currentSpan.SetAttribute("app.message", message);
            return message;
        }

        private async Task<string> DetermineMessage()
        {
//            using var span = _tracer.StartActiveSpan("📖 look up message ✨");
            var db = _redisConnection.GetDatabase();
            var message = "generic hello";
            var result = await db.SetRandomMemberAsync("messages");
            if (result.IsNull)
            {
//                span.AddEvent("message was empty from redis, using default");
            }
            else
            {
                message = result.ToString();
            }
//            span.SetAttribute("app.message", message);
            return message;
        }
    }
}