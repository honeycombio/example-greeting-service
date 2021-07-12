
'use strict';
require("honeycomb-beeline")({
  // Get this via https://ui.honeycomb.io/account after signing up for Honeycomb
  writeKey: "{APIKey}",
  // The name of your app is a good choice to start with
  dataset: "{dataset}",
  serviceName: "node-frontend-service"
});
const express = require('express');

// Constants
const PORT = 7000;
const HOST = 'localhost';

// App
const app = express();
app.get('/', (req, res) => {
  res.send('I am the node frontend service!');
});

app.listen(PORT, HOST);
console.log(`Running node frontend service on http://${HOST}:${PORT}`);