package io.honeycomb.examples.frontend_java;

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
public class NameService {
  @Autowired
	private Tracer tracer;

  private String name_endpoint() {
    String nameEndpointFromEnv = System.getenv().getOrDefault("NAME_ENDPOINT", "http://localhost:8000");
    return nameEndpointFromEnv + "/name";
  }
  @WithSpan
  public String getName() throws URISyntaxException {
    URI name_uri = new URI(name_endpoint());

    HttpClient client = HttpClient.newHttpClient();
    HttpRequest request = HttpRequest.newBuilder()
      .uri(name_uri)
      .header("accept", "application/json")
      .build();
    HttpResponse<String> response = null;
    try {
      Span nameServiceCallSpan = tracer.spanBuilder("✨ call /name ✨").startSpan();
      nameServiceCallSpan.makeCurrent();
      response = client.send(request, HttpResponse.BodyHandlers.ofString());
      nameServiceCallSpan.end();
    } catch (IOException | InterruptedException e) {
      e.printStackTrace();
    }
    return response.body();
  }
}
