import { HoneycombWebSDK } from '@honeycombio/opentelemetry-web';
import { getWebAutoInstrumentations } from '@opentelemetry/auto-instrumentations-web';
const sdk = new HoneycombWebSDK({
    // we're sending data to a local collector here to avoid exposing the API key in the browser
    endpoint: 'http://localhost:55681/v1/traces',
    serviceName: 'egs-browser',
    instrumentations: [getWebAutoInstrumentations()],
});
sdk.start();
