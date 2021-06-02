using System;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;

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
        
        [HttpGet]
        public async Task<string> GetAsync()
        {
            var message = await DetermineMessage();
            return message;
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
