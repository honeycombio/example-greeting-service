require 'opentelemetry/sdk'

# A module to wrap customizations for instrumentation. This represents what
# an organization might do to collect customizations into a library to
# standardize and reuse across multiple services.
module O11yWrapper
  # The BaggageSpanProcessor reads attributes stored in Baggage in the span's
  # parent context and adds them to the span.
  #
  # Add this span processor to the pipeline and then keys and values
  # added to Baggage will appear on subsequent child spans for a trace within
  # this service *and* be propagated to external services via propagation
  # headers. If the external services also have a Baggage span processor, the
  # keys and values will appear in those child spans as well.
  #
  # ⚠ ⚠ ⚠️
  # To repeat: a consequence of adding data to Baggage is that the keys and
  # values will appear in all outgoing HTTP headers from the application.
  # Do not put sensitive information in Baggage.
  class BaggageSpanProcessor < OpenTelemetry::SDK::Trace::SpanProcessor
    def on_start(span, parent_context)
      span.add_attributes(OpenTelemetry::Baggage.values(context: parent_context))
    end
  end

  # CarryOn is not Baggage.
  #
  # Paired with a CarryOnSpanProcessor, adding keys and values to CarryOn will
  # result in those key/values appearing on subsequent child spans for a trace
  # but only within the current process. Data will not automatically propagate
  # outside the process as with Baggage.
  #
  # ... though, you *could* choose to check your CarryOn into Baggage at any
  # point you are comfortable propagating your CarryOn data to other services.
  module CarryOn
    extend self

    # the key under which attributes will be stored for CarryOn within a context
    CONTEXT_KEY = ::OpenTelemetry::Context.create_key('carry-on-attrs')
    private_constant :CONTEXT_KEY

    # retrieve the CarryOn attributes from a given or current context
    def attributes(context = nil)
      context ||= ::OpenTelemetry::Context.current
      context.value(CONTEXT_KEY) || {}
    end

    # return a new Context with the attributes given set within CarryOn
    def with_attributes(attributes_hash)
      attributes_hash = attributes.merge(attributes_hash)
      ::OpenTelemetry::Context.with_value(CONTEXT_KEY, attributes_hash) { |c, h| yield h, c }
    end
  end

  # The CarryOnSpanProcessor reads attributes stored under a custom CarryOn key
  # in the span's parent context and adds them to the span.
  #
  # Add this span processor to the pipeline and then keys and values
  # added to CarryOn will appear on subsequent child spans for a trace only
  # within this service. CarryOn attributes do NOT propagate outside this process.
  class CarryOnSpanProcessor < OpenTelemetry::SDK::Trace::SpanProcessor
    def on_start(span, parent_context)
        span.add_attributes(CarryOn.attributes(parent_context))
    end
  end
end
