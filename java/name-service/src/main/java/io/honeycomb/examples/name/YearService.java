package io.honeycomb.examples.name;

import java.io.IOException;
import java.net.ConnectException;
import java.net.URI;
import java.net.URISyntaxException;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.Tracer;
import io.opentelemetry.extension.annotations.WithSpan;

@Component
public class YearService {
    @Autowired
    private Tracer tracer;

    private String year_endpoint() {
        String yearEndpointFromEnv = "http://" + System.getenv().getOrDefault("YEAR_ENDPOINT", "localhost:6001");
        return yearEndpointFromEnv + "/year";
    }

    @WithSpan
    public String getYear() throws URISyntaxException {
        URI yearUri = new URI(year_endpoint());

        HttpClient client = HttpClient.newHttpClient();
        HttpRequest request = HttpRequest.newBuilder()
            .uri(yearUri)
            .header("accept", "application/json")
            .build();
        HttpResponse<String> response = null;
        Span yearServiceCallSpan = tracer.spanBuilder("✨ call /year ✨").startSpan();
        yearServiceCallSpan.makeCurrent();
        try {
            response = client.send(request, HttpResponse.BodyHandlers.ofString());
        } catch (IOException | InterruptedException e) {
            e.printStackTrace();
        } finally {
            yearServiceCallSpan.end();
        }
        return response == null
            ? ""
            : response.body();
    }
}
