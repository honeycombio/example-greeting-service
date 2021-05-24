package io.honeycomb.examples.name;

import java.io.IOException;
import java.net.URISyntaxException;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import io.opentelemetry.api.trace.Tracer;

@RestController
public class NameController {
	@Autowired
	private Tracer tracer;

	@Autowired
	private NameService nameService;

	@RequestMapping("/name")
	public String index() throws URISyntaxException, IOException, InterruptedException {
		String name = nameService.getName();

		return name;
	}
}
