package io.honeycomb.examples.frontend_java;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;

import io.opentelemetry.api.GlobalOpenTelemetry;
import io.opentelemetry.api.trace.Tracer;
// import io.honeycomb.opentelemetry.HoneycombSdk;
// import io.honeycomb.opentelemetry.sdk.trace.samplers.DeterministicTraceSampler;

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
