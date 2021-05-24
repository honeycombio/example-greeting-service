package io.honeycomb.examples.javaotlp;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;

@RestController
public class YearController {
	@Autowired
	private YearService yearService;

	@RequestMapping("/year")
	public String index() {
    return yearService.getYear();
	}
}
