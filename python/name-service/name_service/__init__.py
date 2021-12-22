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


trace.set_tracer_provider(TracerProvider(
    resource=Resource.create({SERVICE_NAME: "name-python"})
))
tracer = trace.get_tracer_provider().get_tracer(__name__)

trace.get_tracer_provider().add_span_processor(
    BatchSpanProcessor(OTLPSpanExporter(
        headers=(("x-honeycomb-team", os.environ.get("HONEYCOMB_API_KEY")),
                 ("x-honeycomb-dataset", os.environ.get("HONEYCOMB_DATASET"))),
        endpoint=os.environ.get("HONEYCOMB_API_ENDPOINT",
                                "https://api.honeycomb.io")
    )))

app = Flask(__name__)
FlaskInstrumentor().instrument_app(app)
RequestsInstrumentor().instrument()


@app.route('/name')
def get_name():
    year = get_year()
    with tracer.start_as_current_span("📖 look up name based on year ✨"):
        names = names_by_year[year]
    return random.choice(names)
