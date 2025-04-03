package io.honeycomb.examples.frontend_java;

import io.opentelemetry.api.baggage.Baggage;
import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.Tracer;
import io.opentelemetry.context.Scope;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.net.URISyntaxException;

@RestController
public class GreetingController {
    @Autowired
    private NameService nameService;

    @Autowired
    private MessageService messageService;

    @Autowired
    private Tracer tracer;

    @RequestMapping("/greeting")
    public String index() throws URISyntaxException {
        String name = nameService.getName();

        // Simulate random server error
        if (Math.random() < 0.2) { // 20% chance of failure
            throw new RuntimeException("Simulated server error");
        }

        Span.current().setAttribute("app.username", name);

        try (final Scope ignored = Baggage.current()
            .toBuilder()
            .put("app.username", name)
            .build()
            .makeCurrent()
        ) {
            String message = messageService.getMessage();

            Span span = tracer.spanBuilder("🎨 render message ✨").startSpan();
            String greeting = String.format("Hello, %s, %s", name, message);
            span.end();
            return greeting;
        }
    }
}
