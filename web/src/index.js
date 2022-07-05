//TODOS
// Remove local greeting - DONE
// Manually instrument requests
// Connect page load with request

import { trace, context } from '@opentelemetry/api';
// import { tracer } from './tracing-http';
import { determineMessage } from './message';
import { determineYear } from './year';
import { determineName } from './name';

window.onload = (event) => {
  // ... do loading things
  // ... attach timing information
//   documentLoadSpan.end(); //once page is loaded, end the span
};

const createButton = (text, onClick) => {
  const button = document.createElement('button');
  button.textContent = text;
  button.onclick = onClick;
//   button.addEventListener('click', () => {
//     const buttonClickSpan = tracer.startSpan(text, { attributes: { button: text } }, ctx);
//     buttonClickSpan.end();
//   });
  document.body.appendChild(button);
};

const getGreetingContent = async () => {
  let greetingContent;

  try {
    const response = await fetch('http://localhost:7000/greeting');
    if (response.ok) {
      greetingContent = await response.text();
    }
  } catch (error) {
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

main().then(() => {;
});
