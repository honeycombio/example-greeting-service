package io.honeycomb.examples.name;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;

import io.opentelemetry.api.GlobalOpenTelemetry;
import io.opentelemetry.api.trace.Tracer;

@SpringBootApplication
public class NameApplication {

	@Bean
	public Tracer tracer() {
		return GlobalOpenTelemetry.getTracer("frontend-internal");
	}

	public static void main(String[] args) {
		SpringApplication.run(NameApplication.class, args);
	}
}
