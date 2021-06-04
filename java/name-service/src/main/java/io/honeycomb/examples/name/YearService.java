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
    String yearEndpointFromEnv = System.getenv().getOrDefault("YEAR_ENDPOINT", "http://localhost:6001");
    return yearEndpointFromEnv + "/year";
  }

  @WithSpan
  public String getYear() throws URISyntaxException, IOException, InterruptedException {
    URI year_uri = new URI(year_endpoint());

    HttpClient client = HttpClient.newHttpClient();
    HttpRequest request = HttpRequest.newBuilder()
      .uri(year_uri)
      .header("accept", "application/json")
      .build();
    HttpResponse<String> response = null;
    Span year_service_call_span = tracer.spanBuilder("✨ call /year ✨").startSpan();
    year_service_call_span.makeCurrent();
    try {
      response = client.send(request, HttpResponse.BodyHandlers.ofString());
    } catch (IOException | InterruptedException e) {
      e.printStackTrace();
    } finally {
      year_service_call_span.end();
    }
    return response == null
      ? ""
      : response.body();
  }
}
