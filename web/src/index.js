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

const createButton = (text, onClick) => {
  const button = document.createElement('button');
  button.textContent = text;
  button.onclick = onClick;
  button.addEventListener('click', () => {
    const buttonClickSpan = tracer.startSpan(text, { attributes: { button: text } }, ctx);
    buttonClickSpan.end();
  });
  document.body.appendChild(button);
};

const getGreetingContent = async () => {
  let greetingContent;

  try {
    const response = await fetch('http://localhost:7000/greeting');
    if (response.ok) {
      greetingContent = await response.text();
    } else {
      greetingContent = localGreeting();
    }
  } catch (error) {
    greetingContent = localGreeting();
  }

  const greeting =
    document.getElementsByTagName('h1').length === 0
      ? document.createElement('h1')
      : document.getElementsByTagName('h1')[0];
  greeting.innerHTML = greetingContent;

  if (document.getElementsByTagName('h1').length === 0) {
    document.body.appendChild(greeting);
  }
};

const main = async () => {
  createButton('Refresh greeting', getGreetingContent);
};

main().then(() => {
  parentSpan.end();
});
