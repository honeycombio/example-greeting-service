package io.honeycomb.examples.frontend_java;

import io.opentelemetry.api.GlobalOpenTelemetry;
import io.opentelemetry.api.trace.Tracer;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;

import io.opentelemetry.api.GlobalOpenTelemetry;
import io.opentelemetry.api.trace.Tracer;

@SpringBootApplication
public class FrontendApplication {

    @Bean
    public Tracer tracer() {
        return GlobalOpenTelemetry.getTracer("frontend-internal");
    }

    public static void main(String[] args) {
        SpringApplication.run(FrontendApplication.class, args);
    }
}
