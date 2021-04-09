package io.honeycomb.examples.javaotlp;

import java.util.Arrays;

import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.ApplicationContext;
import org.springframework.context.annotation.Bean;

import io.honeycomb.opentelemetry.HoneycombSdk;
import io.honeycomb.opentelemetry.sdk.trace.samplers.DeterministicTraceSampler;

@SpringBootApplication
public class YearApplication {

	@Bean // auto-injects dependencies into your controllers
	public HoneycombSdk getHoneycomb() {
		return HoneycombSdk.builder()
    .setApiKey(System.getenv("HONEYCOMB_WRITE_KEY"))
    .setDataset(System.getenv("HONEYCOMB_DATASET"))
    .setEndpoint("https://api.honeycomb.io")
    .setSampler(new DeterministicTraceSampler(5)) // optional - defaults  to "always sample"
    .build();
	}

	public static void main(String[] args) {
		SpringApplication.run(YearApplication.class, args);
	}
}
