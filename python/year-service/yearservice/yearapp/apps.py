import os
import logging
import beeline
import beeline.propagation.w3c as w3c

from django.apps import AppConfig



class YearappConfig(AppConfig):
    name = 'yearapp'

    def ready(self):
        # If you use uwsgi, gunicorn, celery, or other pre-fork models, see the section below on pre-fork
        # models and do not initialize here.
        beeline.init(
            writekey=os.environ.get("HONEYCOMB_API_KEY"),
            dataset=os.environ.get("HONEYCOMB_DATASET"),
            service_name='year-python',
            debug=True,
            http_trace_parser_hook=w3c.http_trace_parser_hook,
            http_trace_propagation_hook=w3c.http_trace_propagation_hook,
        )
