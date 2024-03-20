package io.honeycomb.examples.javaotlp;

import java.util.HashMap;
import java.util.Map;
import java.util.Random;
import org.apache.logging.log4j.Logger;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.message.ObjectMessage;
import org.springframework.stereotype.Component;
import io.opentelemetry.instrumentation.annotations.WithSpan;
import io.opentelemetry.api.OpenTelemetry;
import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.Tracer;

@Component
public class YearService {
    private static final String[] YEARS = new String[] { "2015", "2016", "2017", "2018", "2019", "2020" };
    private static final Random generator = new Random();

    private Tracer tracer;
    private static final Logger logger = LogManager.getLogger("my-logger");

    public YearService(OpenTelemetry openTelemetry) {
        tracer = openTelemetry.getTracer("year-tracer");
    }

    @WithSpan
    public String getYear() {
        Span span = tracer.spanBuilder("🗓 get-a-year ✨").startSpan();
        int rnd = generator.nextInt(YEARS.length);
        String year = YEARS[rnd];
        Map<String, String> mapMessage = new HashMap<>();
        mapMessage.put("selected_year", year);
        logger.info(new ObjectMessage(mapMessage));
        span.end();

        return year;
    }
}
