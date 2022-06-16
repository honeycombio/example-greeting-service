# Browser greeting service
A browser app that sends OTel traces to a collector that passes it through to Honeycomb. It's possible to run this with a backend to request greetings and trace HTTP requests through to the backend.

## Running the app
```shell
tilt up web
```

This will run the browser app that generates a greeting through JS code in the browser. To run the browser app so that it gets a greeting from a server run
```shell
tilt up web node
```

It should be okay to run with other language backend services too, but they may not be configured with CORS to allow requests from localhost:8080, the node app definitely is so it's safest to use that one.
