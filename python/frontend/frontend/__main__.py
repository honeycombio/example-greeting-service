import os
from frontend import create_app
from werkzeug.serving import run_simple

import beeline
from beeline.middleware.werkzeug import HoneyWSGIMiddleware
import beeline.propagation.w3c as w3c

beeline.init(
    # Get this via https://ui.honeycomb.io/account after signing up for Honeycomb
    writekey=os.environ.get("HONEYCOMB_API_KEY"),
    # The name of your app is a good choice to start with
    api_host=os.environ.get("HONEYCOMB_API_ENDPOINT"),
    service_name='frontend-python',
    debug=True,
    http_trace_parser_hook=w3c.http_trace_parser_hook,
    http_trace_propagation_hook=w3c.http_trace_propagation_hook
)


app = HoneyWSGIMiddleware(create_app())
run_simple('0.0.0.0', 7007, app, use_debugger=True, use_reloader=True)
