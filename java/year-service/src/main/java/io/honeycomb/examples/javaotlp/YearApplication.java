package io.honeycomb.examples.javaotlp;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;

import io.opentelemetry.api.GlobalOpenTelemetry;
import io.opentelemetry.api.trace.Tracer;

@SpringBootApplication
public class YearApplication {

    @Bean
    public Tracer tracer() {
        return GlobalOpenTelemetry.getTracer("year-internal");
    }

    // @Bean
    // public HoneycombSdk honeycomb() {
    // 	return new HoneycombSdk.Builder()
    // 		.setApiKey(System.getenv("HONEYCOMB_API_KEY"))
    // 		.setDataset(System.getenv("HONEYCOMB_DATASET"))
    // 		.setServiceName(System.getenv("SERVICE_NAME"))
    // 		.build();
    // }

    public static void main(String[] args) {
        SpringApplication.run(YearApplication.class, args);
    }
}
