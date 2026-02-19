'use strict';

const { trace } = require('@opentelemetry/api');
const tracer = trace.getTracer('name-service');

const express = require('express');
const fetch = require('node-fetch');

// Constants
const PORT = 8000;
const HOST = '0.0.0.0';
const YEAR_ENDPOINT = process.env.YEAR_ENDPOINT || 'localhost:6001';

const yearURL = `${YEAR_ENDPOINT}/year`;

// App
const app = express();
app.get('/name', async (req, res) => {
  const activeSpan = trace.getActiveSpan();
  if (activeSpan) activeSpan.setAttribute('someContext', 'year');

  const year = await tracer.startActiveSpan('✨ call /year ✨', async (span) => {
    const result = await getYear(yearURL);
    span.end();
    return result;
  });

  return tracer.startActiveSpan('look up name based on year', async (span) => {
    const name = determineName(year);
    span.setAttribute('user_name', name);
    span.end();
    res.send(`${name}`);
  });
});

const names = new Map([
  // prettier-ignore
  [2015, ["sophia", "jackson", "emma", "aiden", "olivia", "liam", "ava", "lucas", "mia", "noah"]],
  // prettier-ignore
  [2016, ["sophia", "jackson", "emma", "aiden", "olivia", "lucas", "ava", "liam", "mia", "noah"]],
  // prettier-ignore
  [2017, ["sophia", "jackson", "olivia", "liam", "emma", "noah", "ava", "aiden", "isabella", "lucas"]],
  // prettier-ignore
  [2018, ["sophia", "jackson", "olivia", "liam", "emma", "noah", "ava", "aiden", "isabella", "caden"]],
  // prettier-ignore
  [2019, ["sophia", "liam", "olivia", "jackson", "emma", "noah", "ava", "aiden", "aria", "grayson"]],
  // prettier-ignore
  [2020, ["olivia", "noah", "emma", "liam", "ava", "elijah", "isabella", "oliver", "sophia", "lucas"]],
]);

const getYear = async (url) =>
  fetch(url)
    .then((data) => {
      return data.text();
    })
    .then((text) => {
      text = Number(text);
      const activeSpan = trace.getActiveSpan();
      if (activeSpan) activeSpan.setAttribute('year', text);
      return text;
    })
    .catch((err) => console.error(`Problem getting year: ${err}`));

const getRandomNumber = (array) => {
  return array[Math.floor(Math.random() * array.length)];
};

const determineName = (year) => {
  console.log(typeof year);
  const namesInYear = names.get(year);
  return getRandomNumber(namesInYear);
};

app.listen(PORT, HOST);
console.log(`Running node name service on http://${HOST}:${PORT}`);
