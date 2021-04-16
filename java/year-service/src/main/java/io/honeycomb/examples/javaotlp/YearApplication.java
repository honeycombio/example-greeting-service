package io.honeycomb.examples.javaotlp; // why do we need this?

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

import io.honeycomb.opentelemetry.HoneycombSdk;

@SpringBootApplication
public class YearApplication {

	HoneycombSdk honeycomb = new HoneycombSdk.Builder().setApiKey(System.getenv("HONEYCOMB_API_KEY"))
			.setDataset(System.getenv("HONEYCOMB_DATASET")).setServiceName("year-service-java").build();

	public static void main(String[] args) {
		SpringApplication.run(YearApplication.class, args);
	}
}
