'use strict';

const { trace } = require('@opentelemetry/api');
const tracer = trace.getTracer('message-service');

const express = require('express');

// Constants
const PORT = 9000;
const HOST = '0.0.0.0';

// App
const app = express();
app.get('/message', async (req, res) => {
  const activeSpan = trace.getActiveSpan();
  if (activeSpan) activeSpan.setAttribute('name', 'Message');

  return tracer.startActiveSpan('look up message', async (span) => {
    const message = determineMessage();
    span.setAttribute('user_message', message);
    span.end();
    res.send(`${message}`);
  });
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
