package io.honeycomb.examples.javaotlp;

import java.util.Random;

import io.opentelemetry.api.OpenTelemetry;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import io.opentelemetry.extension.annotations.WithSpan;

import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.Tracer;

@Component
public class YearService {
    private static final String[] YEARS = new String[]{"2015", "2016", "2017", "2018", "2019", "2020"};
    private static final Random generator = new Random();

    @Autowired
    private OpenTelemetry otel;

    @WithSpan
    public String getYear() {
        Tracer tracer = otel.getTracer("year-internal");
        Span span = tracer.spanBuilder("🗓 get-a-year ✨").startSpan();
        int rnd = generator.nextInt(YEARS.length);
        String year = YEARS[rnd];
        span.end();

        return year;
    }
}
