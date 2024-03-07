import os
from opentelemetry import trace
from opentelemetry.exporter.otlp.proto.grpc.trace_exporter import \
    OTLPSpanExporter
from opentelemetry.sdk.resources import Resource, SERVICE_NAME
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import (
    BatchSpanProcessor,
    ConsoleSpanExporter,
)
from django.apps import AppConfig


class YearappConfig(AppConfig):
    name = 'yearapp'

    def ready(self):
        trace.set_tracer_provider(TracerProvider(
            resource=Resource.create({SERVICE_NAME: "year-python"})
        ))

        # helpful to see in console while developing
        trace.get_tracer_provider().add_span_processor(
            BatchSpanProcessor(ConsoleSpanExporter())
        )

        trace.get_tracer_provider().add_span_processor(
            BatchSpanProcessor(OTLPSpanExporter(
                headers=(
                    ("x-honeycomb-team", os.environ.get("HONEYCOMB_API_KEY")),),
                endpoint=os.environ.get("HONEYCOMB_API_ENDPOINT",
                                        "https://api.honeycomb.io")
            )))
