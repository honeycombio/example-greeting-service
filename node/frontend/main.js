'use strict';
const beeline = require('honeycomb-beeline')({
  // Get this via https://ui.honeycomb.io/account after signing up for Honeycomb
  writeKey: '{APIKey}',
  // The name of your app is a good choice to start with
  dataset: '{dataset}',
  serviceName: 'node-frontend-service',
});

const express = require('express');
const fetch = require('node-fetch');

// Constants
const PORT = 7000;
const HOST = '0.0.0.0';

// TODO - swap out with new endpoints
const nameUrl = 'https://jsonplaceholder.typicode.com/users/1';
const messageUrl = 'https://jsonplaceholder.typicode.com/todos/1';

// App
const app = express();
app.get('/greeting', async (req, res) => {
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
      console.log(json);
      return JSON.stringify(json);
    })
    .catch((err) => console.error('Problem getting name'));

const getMessage = (url) =>
  fetch(url)
    .then((data) => {
      return data.json();
    })
    .then((json) => {
      console.log(json);
      return JSON.stringify(json);
    })
    .catch((err) => console.error('Problem getting message'));

app.listen(PORT, HOST);
console.log(`Running node frontend service on http://${HOST}:${PORT}`);
