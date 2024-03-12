package io.honeycomb.examples.javaotlp;

import io.opentelemetry.api.OpenTelemetry;
import io.opentelemetry.api.GlobalOpenTelemetry;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;

import io.opentelemetry.instrumentation.log4j.appender.v2_17.OpenTelemetryAppender;
import io.opentelemetry.sdk.OpenTelemetrySdk;

@SpringBootApplication
public class YearApplication {

    public static void main(String[] args) {
        OpenTelemetrySdk sdk = OpenTelemetrySdk.builder()
                .build();
        OpenTelemetryAppender.install(sdk);
        SpringApplication.run(YearApplication.class, args);
    }

    @Bean
    public OpenTelemetry openTelemetry() {
        return GlobalOpenTelemetry.get();
    }
}
