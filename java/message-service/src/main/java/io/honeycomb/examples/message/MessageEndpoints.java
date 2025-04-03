package io.honeycomb.examples.message;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.util.MimeTypeUtils;
import org.springframework.web.bind.annotation.*;
import java.util.Random;
import org.springframework.http.HttpStatus;
import org.springframework.web.server.ResponseStatusException;

@RestController
@RequestMapping(value = "message")
public class MessageEndpoints {
    @Autowired
    private MessageService messageService;

    @GetMapping(produces = MimeTypeUtils.APPLICATION_JSON_VALUE)
    public String message() {
        if (new Random().nextDouble() < 0.2) { // 20% chance of failure
            throw new ResponseStatusException(HttpStatus.INTERNAL_SERVER_ERROR, "Simulated server error");
        }
        return messageService.getMessage();
    }
}
