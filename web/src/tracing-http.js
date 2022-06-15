import { trace } from '@opentelemetry/api';
import { ConsoleSpanExporter, SimpleSpanProcessor } from '@opentelemetry/sdk-trace-base';
import { OTLPTraceExporter } from '@opentelemetry/exporter-trace-otlp-http';
import { WebTracerProvider } from '@opentelemetry/sdk-trace-web';
import { ZoneContextManager } from '@opentelemetry/context-zone';
import { Resource } from '@opentelemetry/resources';
import { SemanticResourceAttributes } from '@opentelemetry/semantic-conventions';

const provider = new WebTracerProvider({
  resource: new Resource({
    [SemanticResourceAttributes.SERVICE_NAME]: 'egs-browser',
  }),
});

// Note: For production consider using the "BatchSpanProcessor" to reduce the number of requests
// to your exporter. Using the SimpleSpanProcessor here as it sends the spans immediately to the
// exporter without delay
provider.addSpanProcessor(new SimpleSpanProcessor(new ConsoleSpanExporter()));
provider.addSpanProcessor(new SimpleSpanProcessor(new OTLPTraceExporter({
  url: "http://localhost:55681/v1/traces"
})));
provider.register({
  contextManager: new ZoneContextManager(),
});

const tracer = trace.getTracer();
const root_span = tracer.startSpan('document_load');
//start span when navigating to page
root_span.setAttribute('pageUrlwindow', window.location.href);
window.onload = (event) => {
  // ... do loading things
  // ... attach timing information
 root_span.end(); //once page is loaded, end the span
};
