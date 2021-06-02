using System.Diagnostics;
using OpenTelemetry;

namespace name_service
{
    public class BaggageActivityProcessor : BaseProcessor<Activity>
    {
        public override void OnStart(Activity data)
        {
            foreach (var (key, value) in data.Baggage)
            {
                data.SetTag(key, value);
            }
        }
    }
}
