const { Resource } = require('@opentelemetry/resources');
const { ResourceAttributes } = require('@opentelemetry/semantic-conventions');
// prettier-ignore
const { ExpressInstrumentation } = require('@opentelemetry/instrumentation-express');
const { HttpInstrumentation } = require('@opentelemetry/instrumentation-http');
const { NodeTracerProvider } = require('@opentelemetry/node');
const { registerInstrumentations } = require('@opentelemetry/instrumentation');
const { SimpleSpanProcessor } = require('@opentelemetry/tracing');
// prettier-ignore
const { CollectorTraceExporter } = require('@opentelemetry/exporter-collector-grpc');
const grpc = require('@grpc/grpc-js');
// Honeycomb
const HONEYCOMB_API_KEY = process.env.HONEYCOMB_API_KEY || '';
const HONEYCOMB_DATASET = process.env.HONEYCOMB_DATASET || '';
const SERVICE_NAME = process.env.SERVICE_NAME || 'node-year-service';
const OTLP_ENDPOINT =
  process.env.OTLP_ENDPOINT || 'grpc://api.honeycomb.io:443/';

const provider = new NodeTracerProvider({
  resource: new Resource({
    [ResourceAttributes.SERVICE_NAME]: `${SERVICE_NAME}`,
  }),
});

const metadata = new grpc.Metadata();
metadata.set('x-honeycomb-team', HONEYCOMB_API_KEY);
metadata.set('x-honeycomb-dataset', HONEYCOMB_DATASET);

provider.addSpanProcessor(
  new SimpleSpanProcessor(
    new CollectorTraceExporter({
      url: OTLP_ENDPOINT,
      credentials: grpc.credentials.createSsl(),
      metadata,
    })
  )
);
provider.register();

registerInstrumentations({
  instrumentations: [HttpInstrumentation, ExpressInstrumentation],
});
