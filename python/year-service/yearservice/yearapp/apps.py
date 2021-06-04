from django.apps import AppConfig
import os
import logging
import beeline


class YearappConfig(AppConfig):
    name = 'yearapp'

    def ready(self):
        # If you use uwsgi, gunicorn, celery, or other pre-fork models, see the section below on pre-fork
        # models and do not initialize here.
        beeline.init(
            writekey=os.environ.get("HONEYCOMB_API_KEY"),
            dataset=os.environ.get("HONEYCOMB_DATASET"),
            service_name='year-service-python',
            debug=True,
        )
