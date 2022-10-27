import { ConsoleSpanExporter, SimpleSpanProcessor } from '@opentelemetry/sdk-trace-base';
import { OTLPTraceExporter } from '@opentelemetry/exporter-trace-otlp-http';
import { WebTracerProvider } from '@opentelemetry/sdk-trace-web';
import { ZoneContextManager } from '@opentelemetry/context-zone';
import { Resource } from '@opentelemetry/resources';
import { SemanticResourceAttributes } from '@opentelemetry/semantic-conventions';
import { registerInstrumentations } from '@opentelemetry/instrumentation';
// import { DocumentLoadInstrumentation } from '@opentelemetry/instrumentation-document-load';
// import { XMLHttpRequestInstrumentation } from '@opentelemetry/instrumentation-xml-http-request';
// import { FetchInstrumentation } from '@opentelemetry/instrumentation-fetch';
// import { UserInteractionInstrumentation } from '@opentelemetry/instrumentation-user-interaction';
import { getWebAutoInstrumentations } from '@opentelemetry/auto-instrumentations-web';

const provider = new WebTracerProvider({
  resource: new Resource({
    [SemanticResourceAttributes.SERVICE_NAME]: 'egs-browser',
  }),
});

// Note: For production consider using the "BatchSpanProcessor" to reduce the number of requests
// to your exporter. Using the SimpleSpanProcessor here as it sends the spans immediately to the
// exporter without delay
provider.addSpanProcessor(new SimpleSpanProcessor(new ConsoleSpanExporter()));
provider.addSpanProcessor(
  new SimpleSpanProcessor(
    new OTLPTraceExporter({
      // we're sending data to a local collector here because
      // traces are exported in HTTP/JSON which is not yet
      // natively supported by Honeycomb
      url: 'http://localhost:55681/v1/traces',
    })
  )
);

provider.register({
  contextManager: new ZoneContextManager(),
});

registerInstrumentations({
  instrumentations: [
    getWebAutoInstrumentations(),
    // new DocumentLoadInstrumentation(),
    // new XMLHttpRequestInstrumentation(),
    // new FetchInstrumentation(),
    // new UserInteractionInstrumentation(),
  ],
});
