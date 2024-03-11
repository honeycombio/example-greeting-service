import logging
import os
from opentelemetry import trace
from opentelemetry._logs import set_logger_provider
from opentelemetry.exporter.otlp.proto.grpc.trace_exporter import (
    OTLPSpanExporter
)
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import (
    BatchSpanProcessor,
    ConsoleSpanExporter
)
from opentelemetry.exporter.otlp.proto.grpc._log_exporter import (
    OTLPLogExporter
)
from opentelemetry.sdk._logs import LoggerProvider, LoggingHandler
from opentelemetry.sdk._logs.export import (
    BatchLogRecordProcessor,
    ConsoleLogExporter,
    SimpleLogRecordProcessor
)
from opentelemetry.sdk.resources import (
    Resource,
    SERVICE_NAME
)
from django.apps import AppConfig


class YearappConfig(AppConfig):
    name = 'yearapp'

    def ready(self):
        resource = Resource.create({SERVICE_NAME: 'year-python'})
        headers = (('x-honeycomb-team', os.environ.get('HONEYCOMB_API_KEY')),),
        endpoint = os.environ.get(
            'HONEYCOMB_API_ENDPOINT', 'https://api.honeycomb.io')

        trace.set_tracer_provider(TracerProvider(
            resource=resource
        ))

        trace_exporter = OTLPSpanExporter(
            headers=headers,
            endpoint=endpoint
        )

        trace.get_tracer_provider().add_span_processor(
            BatchSpanProcessor(ConsoleSpanExporter())
        )

        trace.get_tracer_provider().add_span_processor(
            BatchSpanProcessor(trace_exporter))

        logger_provider = LoggerProvider(
            resource=resource
        )
        set_logger_provider(logger_provider)

        log_exporter = OTLPLogExporter(
            headers=headers,
            endpoint=endpoint
        )

        logger_provider.add_log_record_processor(
            SimpleLogRecordProcessor(ConsoleLogExporter()))

        logger_provider.add_log_record_processor(
            BatchLogRecordProcessor(log_exporter)
        )
        logger = logging.getLogger('my-logger')
        logger.addHandler(LoggingHandler(
            level=logging.INFO, logger_provider=logger_provider))

        logger.setLevel(logging.INFO)
