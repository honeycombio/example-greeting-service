'use strict';

const { HoneycombSDK } = require('@honeycombio/opentelemetry-node');
const {
  getNodeAutoInstrumentations,
} = require('@opentelemetry/auto-instrumentations-node');

const sdk = new HoneycombSDK({
  instrumentations: [getNodeAutoInstrumentations()],
  serviceName: "frontend-node",
});

sdk.start()