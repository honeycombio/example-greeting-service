'use strict';
const beeline = require('honeycomb-beeline');

beeline({
  // Get this via https://ui.honeycomb.io/account after signing up for Honeycomb
  writeKey: process.env.HONEYCOMB_API_KEY,
  // The name of your app is a good choice to start with
  dataset: process.env.HONEYCOMB_DATASET,
  serviceName: process.env.SERVICE_NAME || 'node-frontend-service',
  apiHost: process.env.HONEYCOMB_API_ENDPOINT || 'https://api.honeycomb.io',
  httpTraceParserHook: beeline.w3c.httpTraceParserHook,
  httpTracePropagationHook: beeline.w3c.httpTracePropagationHook,
});

const express = require('express');
const fetch = require('node-fetch');
const cors = require('cors');



// Constants
const PORT = 7000;
const HOST = '0.0.0.0';
const MESSAGE_ENDPOINT = process.env.MESSAGE_ENDPOINT || 'localhost:9000';
const NAME_ENDPOINT = process.env.NAME_ENDPOINT || 'localhost:8000';

const nameUrl = `http://${NAME_ENDPOINT}/name`;
const messageUrl = `http://${MESSAGE_ENDPOINT}/message`;

// App
const app = express();

// CORS for use with web frontend
const corsOptions = {
  origin: ['http://localhost:8080'],
  optionsSuccessStatus: 200
};

app.use(cors(corsOptions))

app.get('/greeting', async (req, res) => {
  beeline.addContext({ name: 'Greetings' });
  const greetingSpan = beeline.startSpan({ name: 'Preparing Greeting' });
  beeline.addTraceContext({ name: 'Greetings' });
  beeline.finishSpan(greetingSpan);
  const nameSpan = beeline.startSpan({ name: '✨ call /name ✨' });
  const name = await getName(nameUrl);
  beeline.finishSpan(nameSpan);
  const messageSpan = beeline.startSpan({ name: '✨ call /message ✨' });
  const message = await getMessage(messageUrl);
  beeline.finishSpan(messageSpan);
  res.send(`Hello ${name}, ${message}`);
});

const getName = (url) =>
  fetch(url)
    .then((data) => {
      return data.text();
    })
    .then((text) => {
      console.log(text);
      beeline.addTraceContext({ user_name: text });
      return text;
    })
    .catch((err) => console.error(`Problem getting name: ${err}`));

const getMessage = (url) =>
  fetch(url)
    .then((data) => {
      return data.text();
    })
    .then((text) => {
      console.log(text);
      beeline.addTraceContext({ user_message: text });
      return text;
    })
    .catch((err) => console.error(`Problem getting message: ${err}`));

app.listen(PORT, HOST);
console.log(`Running node frontend service on http://${HOST}:${PORT}`);
