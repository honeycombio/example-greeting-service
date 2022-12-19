package io.honeycomb.examples.name;

import java.io.IOException;
import java.net.URISyntaxException;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class NameController {
	@Autowired
	private NameService nameService;

	@RequestMapping("/name")
	public String index() throws URISyntaxException, IOException, InterruptedException {
        return nameService.getName();
	}
}
