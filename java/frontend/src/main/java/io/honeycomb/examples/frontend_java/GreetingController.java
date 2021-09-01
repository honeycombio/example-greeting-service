package io.honeycomb.examples.frontend_java;

import java.io.IOException;
import java.net.URISyntaxException;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import io.opentelemetry.api.baggage.Baggage;
import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.Tracer;

@RestController
public class GreetingController {
	@Autowired
	private NameService nameService;

	@Autowired
	private MessageService messageService;

	@Autowired
	private Tracer tracer;

	@RequestMapping("/greeting")
	public String index() throws URISyntaxException, IOException, InterruptedException {
		String name = nameService.getName();

		Span.current().setAttribute("app.username", name);
		Baggage.current()
			.toBuilder()
			.put("app.username", name)
			.build()
			.makeCurrent();

		String message = messageService.getMessage();

		Span span = tracer.spanBuilder("🎨 render message ✨").startSpan();
		String greeting = String.format("Hello, %s, %s", name, message);
		span.end();
		return greeting;
	}
}
