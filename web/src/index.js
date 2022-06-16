import { trace, context } from '@opentelemetry/api';
import { tracer } from './tracing-http';
import { determineMessage } from './message';
import { determineYear } from './year';
import { determineName } from './name';

const parentSpan = tracer.startSpan('main');
const ctx = trace.setSpan(context.active(), parentSpan);
const documentLoadSpan = tracer.startSpan('document_load', undefined, ctx);
window.onload = (event) => {
  // ... do loading things
  // ... attach timing information
  documentLoadSpan.end(); //once page is loaded, end the span
};

const fetchGreeting = async () => {
  const response = await fetch('http://localhost:7000/greeting');
  const greeting = await response.text();
  console.log(greeting);
  return greeting;
};

const localGreeting = () => {
  // Get greeting
  const greetingSpan = tracer.startSpan('determineMessage', undefined, ctx);
  const greetingContent = determineMessage();
  greetingSpan.setAttribute('message', greetingContent);
  greetingSpan.end();

  // Get year
  const yearSpan = tracer.startSpan('determineYear', undefined, ctx);
  const year = determineYear();
  yearSpan.setAttribute('year', year);
  yearSpan.end();

  // Get name
  const nameSpan = tracer.startSpan('determineName', undefined, ctx);
  const name = determineName(year);
  nameSpan.setAttribute('chosenName', name);
  nameSpan.end();

  return `Hello ${name}, ${greetingContent}`;
};

const main = async () => {
  const response = await fetch('http://localhost:7000/greeting');
  let greetingContent;
  if (response.ok) {
    greetingContent = await response.text();
  } else {
    greetingContent = localGreeting();
  }

  const greeting = document.createElement('h1');
  greeting.innerHTML = greetingContent;

  document.body.appendChild(greeting);
};

main().then(() => {
  parentSpan.end();
});
