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

messages = [
    "how are you?", "how are you doing?", "what's good?", "what's up?", "how do you do?",
    "sup?", "good day to you", "how are things?", "howzit?", "woohoo",
]
trace.set_tracer_provider(TracerProvider(
    resource=Resource.create({SERVICE_NAME: "message-python"})
))
tracer = trace.get_tracer_provider().get_tracer(__name__)

trace.get_tracer_provider().add_span_processor(
    BatchSpanProcessor(OTLPSpanExporter(
        headers=(("x-honeycomb-team", os.environ.get("HONEYCOMB_API_KEY")),),
        endpoint=os.environ.get("HONEYCOMB_API_ENDPOINT",
                                "https://api.honeycomb.io")
    )))

app = Bottle()
app.wsgi = OpenTelemetryMiddleware(app.wsgi)
RequestsInstrumentor().instrument()

@app.route('/message')
def message():
    with tracer.start_as_current_span("🤖 choosing message✨"):
        return random.choice(messages)

run(app=app, host='0.0.0.0', port=9000)
