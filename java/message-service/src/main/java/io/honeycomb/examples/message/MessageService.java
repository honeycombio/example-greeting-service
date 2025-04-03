package io.honeycomb.examples.message;

import java.util.Random;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.Tracer;
import io.opentelemetry.instrumentation.annotations.WithSpan;

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
        Span messageLookupSpan = tracer.spanBuilder("📖 look up message ✨").startSpan();
        messageLookupSpan.makeCurrent();
        try {
            Thread.sleep(generator.nextInt(1000)); // Simulate 0–999 ms latency
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        }
        int rnd = generator.nextInt(MESSAGES.length);
        String message = MESSAGES[rnd];
        messageLookupSpan.end();
        return message;
    }
}
