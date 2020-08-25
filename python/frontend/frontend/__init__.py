__version__ = '0.1.0'

from werkzeug.routing import Map, Rule
from werkzeug.wrappers import Request, Response
from http.client import HTTPException

from beeline.patch import urllib
import urllib


class Greeting(object):

    def __init__(self):
        self.url_map = Map([
            Rule('/greeting', endpoint='greeting'),
        ])

    def on_greeting(self, request):
        name = self.get_name()
        message = self.get_message()
        return Response('Hello %s, %s' % (name, message), mimetype='text/plain')

    def dispatch_request(self, request):
        adapter = self.url_map.bind_to_environ(request.environ)
        try:
            endpoint, values = adapter.match()
            return getattr(self, 'on_' + endpoint)(request, **values)
        except HTTPException as e:
            return e

    def get_name(self):
        with urllib.request.urlopen('http://localhost:8000/name') as f:
            return f.read().decode('utf-8')

    def get_message(self):
        with urllib.request.urlopen('http://localhost:9000/message') as f:
            return f.read().decode('utf-8')

    def wsgi_app(self, environ, start_response):
        request = Request(environ)
        response = self.dispatch_request(request)
        return response(environ, start_response)

    def __call__(self, environ, start_response):
        return self.wsgi_app(environ, start_response)


def create_app():
    app = Greeting()
    return app
