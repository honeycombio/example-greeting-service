package io.honeycomb.examples.message;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.util.MimeTypeUtils;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping(value = "message")
public class MessageEndpoints {
    @Autowired
    private MessageService messageService;

    @GetMapping(produces = MimeTypeUtils.APPLICATION_JSON_VALUE)
    public String message() {
        return messageService.getMessage();
    }
}
