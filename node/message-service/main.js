'use strict';
const beeline = require('honeycomb-beeline');

beeline({
  // Get this via https://ui.honeycomb.io/account after signing up for Honeycomb
  writeKey: process.env.HONEYCOMB_API_KEY,
  // The name of your app is a good choice to start with
  dataset: process.env.HONEYCOMB_DATASET,
  serviceName: process.env.SERVICE_NAME || 'node-message-service',
  apiHost: process.env.HONEYCOMB_API || 'https://api.honeycomb.io',
  httpTraceParserHook: beeline.w3c.httpTraceParserHook,
  httpTracePropagationHook: beeline.w3c.httpTracePropagationHook,
});

const express = require('express');

// Constants
const PORT = 9000;
const HOST = '0.0.0.0';

// App
const app = express();
app.get('/message', async (req, res) => {
  beeline.addContext({ name: 'Message' });
  const messageSpan = beeline.startSpan({ name: 'look up message' });
  const message = await determineMessage(messages);
  beeline.finishSpan(messageSpan);
  res.send(`${message}`);
});

const messages = [
  'how are you?',
  'how are you doing?',
  "what's good?",
  "what's up?",
  'how do you do?',
  'sup?',
  'good day to you',
  'how are things?',
  'howzit?',
  'woohoo',
];

function determineMessage() {
  return messages[Math.floor(Math.random() * messages.length)];
}

app.listen(PORT, HOST);
console.log(`Running node message service on http://${HOST}:${PORT}`);
