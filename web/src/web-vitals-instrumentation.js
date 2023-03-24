import { onFID, onLCP, onCLS, onINP, onTTFB } from 'web-vitals';
import { InstrumentationBase } from '@opentelemetry/instrumentation';
import { trace, context } from '@opentelemetry/api';
import { hrTime } from '@opentelemetry/core';

export class WebVitalsInstrumentation extends InstrumentationBase {
  onReport(metric, parentSpanContext) {
    const now = hrTime();
    const webVitalsSpan = trace
      .getTracer('web-vitals-instrumentation')
      .startSpan(metric.name, { startTime: now }, parentSpanContext);

    webVitalsSpan.setAttributes({
      [`webvitals.${metric.name}.name`]: metric.name,
      [`webvitals.${metric.name}.id`]: metric.id,
      [`webvitals.${metric.name}.navigationType`]: metric.navigationType,
      [`webvitals.${metric.name}.delta`]: metric.delta,
      [`webvitals.${metric.name}.rating`]: metric.rating,
      [`webvitals.${metric.name}.value`]: metric.value,
      // can expand these into their own attributes!
      [`webvitals.${metric.name}.entries`]: JSON.stringify(metric.entries),
    });
    webVitalsSpan.end();
  }

  enable() {
    if (this.enabled) {
      return;
    }
    this.enabled = true;

    // create a parent span that will have all web vitals spans as children
    const parentSpan = trace.getTracer('web-vitals-instrumentation').startSpan('web-vitals');
    const ctx = trace.setSpan(context.active(), parentSpan);
    parentSpan.end();

    onFID((metric) => {
      this.onReport(metric, ctx);
    });
    onCLS((metric) => {
      this.onReport(metric, ctx);
    });
    onLCP((metric) => {
      this.onReport(metric, ctx);
    });
    onINP((metric) => {
      this.onReport(metric, ctx);
    });
    onTTFB((metric) => {
      this.onReport(metric, ctx);
    });
  }
}
