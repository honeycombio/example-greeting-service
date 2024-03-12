package io.honeycomb.examples.javaotlp;

import io.opentelemetry.api.trace.Tracer;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class YearController {
    @Autowired
    private YearService yearService;

    @RequestMapping("/year")
    public String index() {
        return yearService.getYear();
    }
}
