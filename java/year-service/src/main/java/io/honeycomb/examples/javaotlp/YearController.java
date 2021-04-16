package io.honeycomb.examples.javaotlp;

import org.springframework.web.bind.annotation.RestController;

import java.util.Random;

import org.springframework.web.bind.annotation.RequestMapping;

import io.opentelemetry.api.trace.Span;

@RestController
public class YearController {

	private static final String[] YEARS = new String[]{"2015", "2016", "2017", "2018", "2019", "2020"};
	private static final Random generator = new Random();

	@RequestMapping("/year")
	public String index() {
		Span span = Span.current();

		// import statement stays the same for both vanilla OTel and honey OTel
		span.setAttribute("import", "import io.opentelemetry.api.trace.Span;");

		int rnd = generator.nextInt(YEARS.length);
		String year = YEARS[rnd];
		span.setAttribute("year", year);
		return year;
	}
}
