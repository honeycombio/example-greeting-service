'use strict';
const beeline = require('honeycomb-beeline')({
  // Get this via https://ui.honeycomb.io/account after signing up for Honeycomb
  writeKey: `${process.env.HONEYCOMB_API_KEY}`,
  // The name of your app is a good choice to start with
  dataset: `${process.env.HONEYCOMB_DATASET}`,
  serviceName: `${process.env.SERVICE_NAME}` || 'node-frontend-service',
});

const express = require('express');
const fetch = require('node-fetch');

// Constants
const PORT = 7000;
const HOST = '0.0.0.0';
const MESSAGE_ENDPOINT = process.env.MESSAGE_ENDPOINT || 'localhost:9000';

// TODO - swap out with new endpoints
const nameUrl = 'https://jsonplaceholder.typicode.com/users/1';
const messageUrl = `http://${MESSAGE_ENDPOINT}/message`;

// App
const app = express();
app.get('/greeting', async (req, res) => {
  beeline.addContext({ name: 'Greetings' });
  beeline.addTraceContext({ name: 'Greetings' });
  const name = await getName(nameUrl);
  const message = await getMessage(messageUrl);
  res.send(`Hello ${name}, ${message}`);
});

const getName = (url) =>
  fetch(url)
    .then((data) => {
      return data.json();
    })
    .then((json) => {
      console.log(json.username);
      beeline.addTraceContext({ user_name: json.username });
      return JSON.stringify(json.username);
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
