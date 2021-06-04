package io.honeycomb.examples.name;

import java.io.IOException;
import java.net.URISyntaxException;
import java.util.HashMap;
import java.util.Map;
import java.util.Random;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.Tracer;
import io.opentelemetry.extension.annotations.WithSpan;

@Component
public class NameService {
  @Autowired
  private Tracer tracer;

  private static final Map<Integer,String[]> namesByYear = new HashMap<Integer,String[]>(){{
    this.put(2015, new String[] {"sophia", "jackson", "emma", "aiden", "olivia", "liam", "ava", "lucas", "mia", "noah"});
    this.put(2016, new String[] {"sophia", "jackson", "emma", "aiden", "olivia", "lucas", "ava", "liam", "mia", "noah"});
    this.put(2017, new String[] {"sophia", "jackson", "olivia", "liam", "emma", "noah", "ava", "aiden", "isabella", "lucas"});
    this.put(2018, new String[] {"sophia", "jackson", "olivia", "liam", "emma", "noah", "ava", "aiden", "isabella", "caden"});
    this.put(2019, new String[] {"sophia", "liam", "olivia", "jackson", "emma", "noah", "ava", "aiden", "aira", "grayson"});
    this.put(2020, new String[] {"olivia", "noah", "emma", "liam", "ava", "elijah", "isabella", "oliver", "sophia", "lucas"});
  }};

  private static final Random generator = new Random();

  @Autowired
  YearService yearService;

  @WithSpan
  public String getName() throws NumberFormatException, URISyntaxException, IOException, InterruptedException {
    int year = Integer.parseInt(yearService.getYear());

    Span name_lookup_span = tracer.spanBuilder("📖 look up name based on year ✨").startSpan();
    name_lookup_span.makeCurrent();
    String[] candidate_names = namesByYear.get(year);
    int rnd = generator.nextInt(candidate_names.length);
    String name = candidate_names[rnd];
    name_lookup_span.end();

    return name;
  }
}
