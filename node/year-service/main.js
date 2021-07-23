'use strict';
const grpc = require('@grpc/grpc-js');
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
const opentelemetry = require('@opentelemetry/api');

const express = require('express');

// Constants
const PORT = 6001;
const HOST = '0.0.0.0';

// Honeycomb
const HONEYCOMB_API_KEY = process.env.HONEYCOMB_API_KEY || '';
const HONEYCOMB_DATASET = process.env.HONEYCOMB_DATASET || '';
const SERVICE_NAME = process.env.SERVICE_NAME || 'node-year-service';

// App
const app = express();
app.get('/year', async (req, res) => {
  const span = opentelemetry.trace
    .getTracer('default')
    .startSpan('Getting year');
  const year = await determineYear(years);
  res.send(`${year}`);
  span.end();
});

const years = [2015, 2016, 2017, 2018, 2019, 2020];

function determineYear() {
  return years[Math.floor(Math.random() * years.length)];
}

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
      url: 'grpc://api.honeycomb.io:443/',
      credentials: grpc.credentials.createSsl(),
      metadata,
    })
  )
);
provider.register();

registerInstrumentations({
  instrumentations: [HttpInstrumentation, ExpressInstrumentation],
});

app.listen(PORT, HOST);
console.log(`Running node year service on http://${HOST}:${PORT}`);
