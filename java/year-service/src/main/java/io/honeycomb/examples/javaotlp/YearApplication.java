package io.honeycomb.examples.javaotlp;

import io.opentelemetry.api.OpenTelemetry;
import io.opentelemetry.api.GlobalOpenTelemetry;
import io.opentelemetry.api.trace.Tracer;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;

@SpringBootApplication
public class YearApplication {

    public static void main(String[] args) {
        SpringApplication.run(YearApplication.class, args);
    }

    @Bean
    public OpenTelemetry openTelemetry() {
        return GlobalOpenTelemetry.get();
  }
}
