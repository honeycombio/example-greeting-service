using System.Diagnostics;
using OpenTelemetry;

namespace name_service
{
    public class BaggageSpanProcessor : BaseProcessor<Activity>
    {
        public override void OnEnd(Activity activity)
        {
            foreach (var (key, value) in Baggage.Current)
            {
                activity.SetTag(key, value);
            }
        }
    }
}
