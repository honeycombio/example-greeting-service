import os
import random

from bottle import Bottle, run
from opentelemetry import trace
from opentelemetry.exporter.otlp.proto.grpc.trace_exporter import \
    OTLPSpanExporter
from opentelemetry.instrumentation.requests import RequestsInstrumentor
from opentelemetry.instrumentation.wsgi import OpenTelemetryMiddleware
from opentelemetry.sdk.resources import Resource, SERVICE_NAME
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import BatchSpanProcessor
from honeycomb.opentelemetry import configure_opentelemetry, HoneycombOptions

configure_opentelemetry(
    HoneycombOptions(
        debug=True,
        apikey=os.getenv("HONEYCOMB_API_KEY"),
        service_name="otel-python-example"
    )
)

messages = [
    "how are you?", "how are you doing?", "what's good?", "what's up?", "how do you do?",
    "sup?", "good day to you", "how are things?", "howzit?", "woohoo",
]

tracer = trace.get_tracer(__name__)

app = Bottle()
app.wsgi = OpenTelemetryMiddleware(app.wsgi)
RequestsInstrumentor().instrument()

@app.route('/message')
def message():
    with tracer.start_as_current_span("🤖 choosing message✨"):
        return random.choice(messages)

run(app=app, host='0.0.0.0', port=9000)
