'use strict';
const opentelemetry = require('@opentelemetry/api');

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

const tracer = opentelemetry.trace.getTracer(
  'default'
);

app.get('/greeting', async (req, res) => {
  const greetingSpan = tracer.startSpan('✨ preparing greeting ✨');
  greetingSpan.end()

  const nameSpan = tracer.startSpan('✨ call /name ✨');
  const name = await getName(nameUrl);
  nameSpan.setAttribute("app.user_name", name);
  nameSpan.end()

  const messageSpan = tracer.startSpan('✨ call /message ✨');
  const message = await getMessage(messageUrl);
  messageSpan.end()

  const responseSpan = tracer.startSpan('✨ post response ✨');
  res.send(`Hello ${name}, ${message}`);
  responseSpan.end()
});

const getName = (url) =>
  fetch(url)
    .then((data) => {
      return data.text();
    })
    .then((text) => {
      console.log(text);
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
      return text;
    })
    .catch((err) => console.error(`Problem getting message: ${err}`));

app.listen(PORT, HOST);
console.log(`Running node frontend service on http://${HOST}:${PORT}`);
