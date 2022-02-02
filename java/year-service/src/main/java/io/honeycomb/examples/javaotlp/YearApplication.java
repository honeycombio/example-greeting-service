package io.honeycomb.examples.javaotlp;

import io.honeycomb.opentelemetry.OpenTelemetryConfiguration;
import io.honeycomb.opentelemetry.sdk.trace.spanprocessors.BaggageSpanProcessor;
import io.opentelemetry.api.OpenTelemetry;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;

@SpringBootApplication
public class YearApplication {

    @Bean
    public OpenTelemetry honeycomb() {
        return OpenTelemetryConfiguration.builder()
            .addSpanProcessor(new BaggageSpanProcessor())
            .setApiKey(System.getenv("HONEYCOMB_API_KEY"))
            .setDataset(System.getenv("HONEYCOMB_DATASET"))
            .setServiceName(System.getenv("SERVICE_NAME"))
            .setEndpoint(System.getenv("HONEYCOMB_API_ENDPOINT"))
            .buildAndRegisterGlobal();
    }

    public static void main(String[] args) {
        SpringApplication.run(YearApplication.class, args);
    }
}
