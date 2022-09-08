const process = require('process');
const { Metadata, credentials } = require('@grpc/grpc-js');

const { NodeSDK } = require('@opentelemetry/sdk-node');
const { getNodeAutoInstrumentations } = require('@opentelemetry/auto-instrumentations-node');
const { Resource } = require('@opentelemetry/resources');
const { SemanticResourceAttributes } = require('@opentelemetry/semantic-conventions');
const { OTLPTraceExporter } = require('@opentelemetry/exporter-trace-otlp-grpc');


// Honeycomb
const HONEYCOMB_API_KEY = process.env.HONEYCOMB_API_KEY || '';
const HONEYCOMB_DATASET = process.env.HONEYCOMB_DATASET || '';
const SERVICE_NAME = process.env.SERVICE_NAME || 'node-frontend-service';
const OTLP_ENDPOINT = process.env.HONEYCOMB_API_ENDPOINT || 'grpc://api.honeycomb.io:443/';

const metadata = new Metadata();
metadata.set('x-honeycomb-team', HONEYCOMB_API_KEY);
metadata.set('x-honeycomb-dataset', HONEYCOMB_DATASET);

const traceExporter = new OTLPTraceExporter({
  url: OTLP_ENDPOINT,
  credentials: credentials.createSsl(),
  metadata,
});
const sdk = new NodeSDK({
  resource: new Resource({
    [SemanticResourceAttributes.SERVICE_NAME]: SERVICE_NAME,
    [SemanticResourceAttributes.SERVICE_VERSION]: "0.1.0",
  }),
  traceExporter,
  instrumentations: [getNodeAutoInstrumentations()],
});

sdk
  .start()
  .then(() => console.log('Tracing initialized'))
  .catch((error) => console.log('Error initializing tracing', error));

process.on('SIGTERM', () => {
  sdk
    .shutdown()
    .then(() => console.log('Tracing terminated'))
    .catch((error) => console.log('Error terminating tracing', error))
    .finally(() => process.exit(0));
});
