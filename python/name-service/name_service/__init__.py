__version__ = '0.1.0'

import logging
import os
import random

import requests
from flask import Flask
from opentelemetry import trace

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

tracer = trace.get_tracer(__name__)

logger = logging.getLogger("my-logger")
logger.setLevel(logging.INFO)

app = Flask(__name__)

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
