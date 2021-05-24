package io.honeycomb.examples.message;

import java.util.Random;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.Tracer;
import io.opentelemetry.extension.annotations.WithSpan;

@Component
public class MessageService {
    @Autowired
	private Tracer tracer;

    private static final String[] MESSAGES = new String[] { "how are you?", "how are you doing?", "what's good?", "what's up?", "how do you do?",
    "sup?", "good day to you", "how are things?", "howzit?", "woohoo" };
    private static final Random generator = new Random();

    @WithSpan
    public String getMessage() {
        return pickMessage();
    }

    private String pickMessage() {
        Span message_lookup_span = tracer.spanBuilder("📖 look up message ✨").startSpan();
        message_lookup_span.makeCurrent();
        int rnd = generator.nextInt(MESSAGES.length);
        String message = MESSAGES[rnd];
        message_lookup_span.end();
        return message;
    }
}
