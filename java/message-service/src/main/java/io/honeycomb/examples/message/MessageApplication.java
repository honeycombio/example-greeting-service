package io.honeycomb.examples.message;

import io.opentelemetry.api.GlobalOpenTelemetry;
import io.opentelemetry.api.trace.Tracer;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;

@SpringBootApplication
public class MessageApplication {
    @Bean
	public Tracer tracer() {
		return GlobalOpenTelemetry.getTracer("frontend-internal");
	}

    public static void main(String[] args) {
        SpringApplication.run(MessageApplication.class, args);
    }
}
