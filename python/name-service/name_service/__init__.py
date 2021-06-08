__version__ = '0.1.0'

import os
import random

import beeline
import beeline.propagation.w3c as w3c
import requests
from beeline.middleware.flask import HoneyMiddleware
from beeline.patch.requests import *
from flask import Flask

names_by_year = {
    2015: ['sophia', 'jackson', 'emma', 'aiden', 'olivia', 'liam', 'ava', 'lucas', 'mia', 'noah'],
    2016: ['sophia', 'jackson', 'emma', 'aiden', 'olivia', 'lucas', 'ava', 'liam', 'mia', 'noah'],
    2017: ['sophia', 'jackson', 'olivia', 'liam', 'emma', 'noah', 'ava', 'aiden', 'isabella', 'lucas'],
    2018: ['sophia', 'jackson', 'olivia', 'liam', 'emma', 'noah', 'ava', 'aiden', 'isabella', 'caden'],
    2019: ['sophia', 'liam', 'olivia', 'jackson', 'emma', 'noah', 'ava', 'aiden', 'aira', 'grayson'],
    2020: ['olivia', 'noah', 'emma', 'liam', 'ava', 'elijah', 'isabella', 'oliver', 'sophia', 'lucas']
}

YEAR_ENDPOINT = os.environ.get('YEAR_ENDPOINT', 'http://localhost:6001') + '/year'

@beeline.traced(name="✨ call /year ✨")
def get_year():
    r = requests.get(YEAR_ENDPOINT)
    return int(r.text)


app = Flask(__name__)
beeline.init(
    # Get this via https://ui.honeycomb.io/account after signing up for Honeycomb
    writekey=os.environ.get("HONEYCOMB_API_KEY"),
    # The name of your app is a good choice to start with
    dataset=os.environ.get("HONEYCOMB_DATASET"),
    service_name='name-service-python',
    debug=True,
    http_trace_parser_hook=w3c.http_trace_parser_hook,
    http_trace_propagation_hook=w3c.http_trace_propagation_hook
)
HoneyMiddleware(app, db_events=True)


@app.route('/name')
def get_name():
    year = get_year()
    with beeline.tracer(name="📖 look up name based on year ✨"):
        names = names_by_year[year]
    return random.choice(names)
