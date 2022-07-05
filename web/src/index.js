//TODOS
// Manually instrument requests
    // parent span
    // request function
    // traceparent header
// Connect page load with request

import { trace, context, SpanStatusCode } from '@opentelemetry/api';

window.onload = (event) => {
    trace.getTracer('event-tracer').startActiveSpan(`Event: ${event.type}`, (span) => {
        span.setAttributes({
            duration_ms: event.timeStamp
        });
        span.end();
    });
};

const request = async (url, method = "GET", headers, body) => {
    return trace.getTracer('request-tracer').startActiveSpan(`Request: ${method} ${url}`, async span => {
        // construct traceparent header 
        const traceparent = `00-${span.spanContext().traceId}-${span.spanContext().spanId}-01`;

        try {
            const response = await fetch(url, {
                method,
                headers: { ...headers, traceparent},
                body
            });

            if (response.ok && response.status >= 200 && response.status < 400) {
                span.setStatus({ code: SpanStatusCode.OK })
                return response.text()
            } else {
                throw new Error(`Request Error ${response.status} ${response.statusText}`)
            }
        
        } catch (error) {
            span.setStatus({ code: SpanStatusCode.ERROR, message: error.message });
        } finally {
            span.end()
        }
    });
}

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
    try {
        const greetingContent = await request('http://localhost:7000/greeting')

        const greeting =
        document.getElementsByTagName('h1').length === 0
          ? document.createElement('h1')
          : document.getElementsByTagName('h1')[0];
      greeting.innerHTML = greetingContent;
    
      if (document.getElementsByTagName('h1').length === 0) {
        document.body.appendChild(greeting);
      }

    } catch (error) {

    }
};

const main = async () => {
  createButton('Refresh greeting', getGreetingContent);
};

main().then(() => {;
});
