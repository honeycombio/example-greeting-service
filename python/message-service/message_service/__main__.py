import os
import beeline
import libhoney
import random
from bottle import Bottle, run
from beeline.middleware.bottle import HoneyWSGIMiddleware

beeline.init(
    # Get this via https://ui.honeycomb.io/account after signing up for Honeycomb
    writekey=os.environ.get("HONEYCOMB_WRITE_KEY"),
    # The name of your app is a good choice to start with
    dataset=os.environ.get("HONEYCOMB_DATASET"),
    service_name='message-service-python',
    debug=True,
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

run(app=app, host='127.0.0.1', port=9000)
