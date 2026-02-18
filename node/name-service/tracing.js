'use strict';

// @opentelemetry/configuration enables OTEL_EXPERIMENTAL_CONFIG_FILE support.
// Full integration with NodeSDK is in progress; trace/log export currently
// uses OTEL_* env vars via NodeSDK's existing env-var config path.
require('@opentelemetry/configuration');

const { NodeSDK } = require('@opentelemetry/sdk-node');
const { getNodeAutoInstrumentations } = require('@opentelemetry/auto-instrumentations-node');

const sdk = new NodeSDK({
  instrumentations: [getNodeAutoInstrumentations({
    '@opentelemetry/instrumentation-fs': { enabled: false },
  })],
});

sdk.start();

process.on('SIGTERM', () => {
  sdk
    .shutdown()
    .then(() => console.log('Tracing terminated'))
    .catch((error) => console.log('Error terminating tracing', error))
    .finally(() => process.exit(0));
});
