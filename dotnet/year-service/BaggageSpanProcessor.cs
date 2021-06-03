using System.Diagnostics;
using OpenTelemetry;

namespace year_service
{
    public class BaggageSpanProcessor : BaseProcessor<Activity>
    {
        public override void OnStart(Activity activity)
        {
            foreach (var (key, value) in Baggage.Current)
            {
                activity.SetTag(key, value);
            }
        }
    }
}
