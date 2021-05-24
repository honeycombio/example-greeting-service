package io.honeycomb.examples.name;

import java.io.IOException;
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
public class MessageService {
  @Autowired
	private Tracer tracer;

  private String message_endpoint = "http://localhost:9000/message";

  @WithSpan
  public String getMessage() throws URISyntaxException, IOException, InterruptedException {
    URI message_uri = new URI(message_endpoint);

    HttpClient client = HttpClient.newHttpClient();
    HttpRequest request = HttpRequest.newBuilder()
        .uri(message_uri)
        .header("accept", "application/json")
        .build();
    HttpResponse<String> response = null;
    try {
      Span message_service_call_span = tracer.spanBuilder("✨ call /name ✨").startSpan();
      message_service_call_span.makeCurrent();
      response = client.send(request, HttpResponse.BodyHandlers.ofString());
      message_service_call_span.end();
    } catch (IOException | InterruptedException e) {
        e.printStackTrace();

      }
    return response.body();
  }
}
