import { trace, SpanStatusCode } from '@opentelemetry/api';

// General request handler, instrumented with OTel
// Forwards traceparent header to connect spans created in the browser
// with spans created on the backend
const request = async (url, method = 'GET', headers, body) => {
  return trace
    .getTracer('request-tracer')
    .startActiveSpan(`Request: ${method} ${url}`, async (span) => {
      // construct traceparent header
      const traceparent = `00-${span.spanContext().traceId}-${span.spanContext().spanId}-01`;

      try {
        const response = await fetch(url, {
          method,
          headers: { ...headers, traceparent },
          body,
        });
        span.setAttributes({
          'http.method': method,
          'http.url': url,
          'response.status_code': response.status,
        });
        if (response.ok && response.status >= 200 && response.status < 400) {
          span.setStatus({ code: SpanStatusCode.OK });
          return response.text();
        } else {
          throw new Error(`Request Error ${response.status} ${response.statusText}`);
        }
      } catch (error) {
        span.setStatus({ code: SpanStatusCode.ERROR, message: error.message });
        throw new Error(error);
      } finally {
        span.end();
      }
    });
};

// Generic button creator function, automatically instruments
// onclick handler to link button clicks to whatever action is taken
const createButton = (text, onClick) => {
  const button = document.createElement('button');
  button.textContent = text;
  button.onclick = () =>
    trace
      .getTracer('button-onclick-tracer')
      .startActiveSpan(`Event: Button Click ${text}`, (span) => {
        onClick();
        span.end();
      });
  document.body.appendChild(button);
};

const updateGreetingContent = async () => {
  const greeting =
    document.getElementsByTagName('h1').length === 0
      ? document.createElement('h1')
      : document.getElementsByTagName('h1')[0];
  try {
    const greetingContent = await request(FRONTEND_ENDPOINT + '/greeting');

    greeting.innerHTML = greetingContent;

    if (document.getElementsByTagName('h1').length === 0) {
      document.body.appendChild(greeting);
    }
  } catch (error) {
    greeting.innerHTML = error.message;

    if (document.getElementsByTagName('h1').length === 0) {
      document.body.appendChild(greeting);
    }
  }
};

const main = async () => {
  await updateGreetingContent();
  createButton('Refresh greeting', updateGreetingContent);
  const link = document.createElement('a');
  link.href = 'https://ui.honeycomb.io';
  link.innerHTML = 'click me!';
  document.body.appendChild(link);
};

// onload kicks off the main function and instruments the page load
// to link it to requests fired off as a result of the page load
window.onload = (event) => {
  trace.getTracer('event-tracer').startActiveSpan(`Event: ${event.type}`, (span) => {
    span.setAttributes({
      duration_ms: event.timeStamp,
    });
    main();
    span.end();
  });
};
