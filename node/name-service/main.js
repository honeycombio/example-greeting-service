'use strict';
const beeline = require('honeycomb-beeline')({
  // Get this via https://ui.honeycomb.io/account after signing up for Honeycomb
  writeKey: `${process.env.HONEYCOMB_API_KEY}`,
  // The name of your app is a good choice to start with
  dataset: `${process.env.HONEYCOMB_DATASET}`,
  serviceName: `${process.env.SERVICE_NAME}` || 'node-name-service',
  apiHost: `${process.env.HONEYCOMB_API}` || 'https://api.honeycomb.io',
});

const express = require('express');
const fetch = require('node-fetch');

// Constants
const PORT = 8000;
const HOST = '0.0.0.0';
const YEAR_ENDPOINT = process.env.YEAR_ENDPOINT || 'localhost:6001';

const yearURL = `http://${YEAR_ENDPOINT}/year`;

// App
const app = express();
app.get('/name', async (req, res) => {
  const name = await determineName();
  const promises = [name, getYear, determineName];
  beeline.addContext({ name: 'Name' });
  Promise.all(promises).then((results) => console.log(results));
  res.send(`${name}`);
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

const getYear = (url) =>
  fetch(url)
    .then((data) => {
      console.log(data);
      console.log(data.text());
      return data.text();
    })
    .then((text) => {
      console.log(text);
      beeline.addTraceContext({ year: text });
      return text;
    })
    .catch((err) => console.error(`Problem getting year: ${err}`));

const getRandomNumber = (array) => {
  return array[Math.floor(Math.random() * array.length)];
};

const determineName = async () => {
  const year = await getYear(yearURL);
  // const year = 2020;
  console.log(`year: ${year}`);
  console.log(typeof year);
  const namesInYear = names.get(year);
  console.log(`namesInYear: ${namesInYear}`);
  console.log(`getRandomNumber(namesInYear): ${getRandomNumber(namesInYear)}`);
  return getRandomNumber(namesInYear);
}

app.listen(PORT, HOST);
console.log(`Running node name service on http://${HOST}:${PORT}`);
