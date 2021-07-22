'use strict';
const beeline = require('honeycomb-beeline')({
  // Get this via https://ui.honeycomb.io/account after signing up for Honeycomb
  writeKey: '{APIKey}',
  // The name of your app is a good choice to start with
  dataset: '{dataset}',
  serviceName: 'node-message-service',
});

const express = require('express');

// Constants
const PORT = 9000;
const HOST = '0.0.0.0';

// App
const app = express();
app.get('/message', async (req, res) => {
  beeline.addContext({ name: 'Message' });
  const message = await determineMessage(messages);
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
