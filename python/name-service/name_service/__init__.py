__version__ = '0.1.0'

import os
import random

import requests
from flask import Flask
from opentelemetry import trace
from opentelemetry.exporter.otlp.proto.grpc.trace_exporter import \
    OTLPSpanExporter
from opentelemetry.instrumentation.flask import FlaskInstrumentor
from opentelemetry.instrumentation.requests import RequestsInstrumentor
from opentelemetry.sdk.resources import Resource, SERVICE_NAME
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import BatchSpanProcessor


import logging
from opentelemetry.exporter.otlp.proto.grpc._log_exporter import (
    OTLPLogExporter,
)
from opentelemetry._logs import set_logger_provider
from opentelemetry.sdk._logs import (
    LoggerProvider,
    LoggingHandler,
)
from opentelemetry.sdk._logs.export import BatchLogRecordProcessor

names_by_year = {
    2015: ['sophia', 'jackson', 'emma', 'aiden', 'olivia', 'liam', 'ava',
           'lucas', 'mia', 'noah'],
    2016: ['sophia', 'jackson', 'emma', 'aiden', 'olivia', 'lucas', 'ava',
           'liam', 'mia', 'noah'],
    2017: ['sophia', 'jackson', 'olivia', 'liam', 'emma', 'noah', 'ava',
           'aiden', 'isabella', 'lucas'],
    2018: ['sophia', 'jackson', 'olivia', 'liam', 'emma', 'noah', 'ava',
           'aiden', 'isabella', 'caden'],
    2019: ['sophia', 'liam', 'olivia', 'jackson', 'emma', 'noah', 'ava',
           'aiden', 'aira', 'grayson'],
    2020: ['olivia', 'noah', 'emma', 'liam', 'ava', 'elijah', 'isabella',
           'oliver', 'sophia', 'lucas']
}

YEAR_ENDPOINT = os.environ.get('YEAR_ENDPOINT',
                               'http://localhost:6001') + '/year'


def get_year():
    with tracer.start_as_current_span("✨ call /year ✨"):
        r = requests.get(YEAR_ENDPOINT)
        return int(r.text)

# shared resource service name
myResource = Resource.create({"service.name": "name-python"})

# tracing pipeline
trace.set_tracer_provider(TracerProvider(resource=myResource))
tracer = trace.get_tracer_provider().get_tracer(__name__)
trace_exporter = OTLPSpanExporter(
        headers=(("x-honeycomb-team", os.environ.get("HONEYCOMB_API_KEY")),),
        endpoint=os.environ.get("HONEYCOMB_API_ENDPOINT",
                                "https://api.honeycomb.io")
    )
trace.get_tracer_provider().add_span_processor(BatchSpanProcessor(trace_exporter))

# logging pipeline
logger_provider = LoggerProvider(resource=myResource)
set_logger_provider(logger_provider)
log_exporter = OTLPLogExporter(
        headers=(("x-honeycomb-team", os.environ.get("HONEYCOMB_API_KEY")),),
        endpoint=os.environ.get("HONEYCOMB_API_ENDPOINT",
                                "https://api.honeycomb.io")
   )
logger_provider.add_log_record_processor(BatchLogRecordProcessor(log_exporter))
handler = LoggingHandler(level=logging.NOTSET, logger_provider=logger_provider)
logging.getLogger().addHandler(handler)
logger = logging.getLogger("my-logger")
logger.setLevel(logging.INFO)

app = Flask(__name__)
FlaskInstrumentor().instrument_app(app)
RequestsInstrumentor().instrument()


@app.route('/name')
def get_name():
    year = get_year()
    current_span = trace.get_current_span()
    logger.info("Selected year: %s", year)
    current_span.set_attribute("app.year_selected", year)

    with tracer.start_as_current_span("📖 look up name based on year ✨"):
        span = trace.get_current_span()
        names = names_by_year[year]
        name = random.choice(names)
        span.set_attribute("app.name_selected", name)
    return name
