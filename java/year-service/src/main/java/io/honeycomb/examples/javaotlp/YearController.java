package io.honeycomb.examples.javaotlp;

import org.springframework.web.bind.annotation.RestController;

import java.util.Random;

import org.springframework.web.bind.annotation.RequestMapping;

import io.opentelemetry.api.trace.Span;

@RestController
public class YearController {

	private static final String[] YEARS = new String[]{"2015", "2016", "2017", "2018", "2019", "2020"};
	private static final Random generator = new Random();

	@RequestMapping("/")
	public String index() {
			// access the current context and add a custom attribute
			Span span = Span.current();
			span.setAttribute("custom_field", "important value");
			return "hello world";
	}
}
