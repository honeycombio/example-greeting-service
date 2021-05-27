using System;
using Microsoft.AspNetCore.Mvc;

namespace year_service.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class YearController : ControllerBase
    {
        private static readonly int[] Years = {
            2015, 2016, 2017, 2018, 2019, 2020
        };


        [HttpGet]
        public int Get()
        {
            var rng = new Random();
            var i = rng.Next(0, 5);
            return Years[i];
        }
    }
}