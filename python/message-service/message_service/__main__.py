import os
import beeline
import beeline.propagation.w3c as w3c
import libhoney
import random
from bottle import Bottle, run
from beeline.middleware.bottle import HoneyWSGIMiddleware

beeline.init(
    # Get this via https://ui.honeycomb.io/account after signing up for Honeycomb
    writekey=os.environ.get("HONEYCOMB_API_KEY"),
    # The name of your app is a good choice to start with
    dataset=os.environ.get("HONEYCOMB_DATASET"),
    service_name='message-service-python',
    debug=True,
    http_trace_parser_hook=w3c.http_trace_parser_hook,
    http_trace_propagation_hook=w3c.http_trace_propagation_hook,
)


messages = [
    "how are you?", "how are you doing?", "what's good?", "what's up?", "how do you do?",
    "sup?", "good day to you", "how are things?", "howzit?", "woohoo",
]


app = Bottle()


@app.route('/message')
def message():
    return random.choice(messages)


app = HoneyWSGIMiddleware(app)

run(app=app, host='0.0.0.0', port=9000)
