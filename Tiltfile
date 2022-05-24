# prints show up in the Tiltfile log so you know what's been run
print("hello my friends")

docker_compose("./docker-compose.yml")

local_resource(
  'curl greeting',
  cmd='curl localhost:7000/greeting',
  trigger_mode=TRIGGER_MODE_MANUAL,
  auto_init=False)

def launch_go_svc(name, dirname="", flags="", auto_init=True):
    '''
    Starts a single go service.

    Parameters:
    name: used to display the name of the process in the tilt tab
    dirname: (optional) directory name in which to run `go run main.go` defaults to 'name'
    flags: (optional) any additional flags to add to the command line
    '''

    cmd = "cd {} && go run main.go -debug {}".format(
        dirname if dirname else name,
        flags if flags else ""
    )
    env = {
        'NAME_ENDPOINT': 'http://localhost:8000',
        'YEAR_ENDPOINT': 'http://localhost:6001',
        'MESSAGE_ENDPOINT': 'http://localhost:9000',
    }
    print("About to start {} with command {}".format(name, cmd))
    local_resource(name, "", auto_init=auto_init, serve_cmd=cmd, serve_env=env)


def launch_go_frontend(auto_init=True):
    launch_go_svc("frontend-go", dirname="golang/frontend")


def launch_go_message_service(auto_init=True):
    launch_go_svc("message-go", dirname="golang/message-service")


def launch_go_name_service(auto_init=True):
    launch_go_svc("name-go", dirname="golang/name-service")


def launch_go_year_service(auto_init=True):
    launch_go_svc("year-go", dirname="golang/year-service")


def launch_python_frontend(auto_init=True):
    cmd = "cd python/frontend && poetry install --no-root && poetry run python -m frontend"
    local_resource("frontend-py", "", auto_init=auto_init, serve_cmd=cmd)


def launch_python_message_service(auto_init=True):
    cmd = "cd python/message-service && poetry install --no-root && poetry run python -m message_service"
    local_resource("message-py", "", auto_init=auto_init, serve_cmd=cmd)


def launch_python_name_service(auto_init=True):
    cmd = "cd python/name-service && poetry install --no-root && poetry run flask run"
    local_resource("name-py", "", auto_init=auto_init, serve_cmd=cmd)


def launch_python_year_service(auto_init=True):
    cmd = "cd python/year-service && poetry install --no-root && poetry run yearservice/manage.py runserver 127.0.0.1:6001"
    local_resource("year-py", "", auto_init=auto_init, serve_cmd=cmd)

def launch_ruby_svc(name, dirname, run_cmd, auto_init=True):
    '''
    Starts a single Ruby service.

    Parameters:
    name: used to display the name of the process in the tilt tab
    dirname: (optional) directory name in which to run `go run main.go` defaults to 'name'
    flags: (optional) any additional flags to add to the command line
    '''

    env = {
        'SERVICE_NAME': name,
        'BUNDLE_BIN': "./.direnv/bin",
        'GEM_HOME': "./.direnv/ruby",
        'OTEL_EXPORTER_OTLP_ENDPOINT': "http://localhost:55681",
    }
    setup_cmd = "cd {} && bundle install".format(dirname)
    serve_cmd = "cd {} && bundle exec {}".format(dirname,run_cmd)
    print("About to start {} with command {}".format(name, serve_cmd))
    local_resource(name, setup_cmd, env=env, auto_init=auto_init, serve_cmd=serve_cmd, serve_env=env)

def launch_ruby_frontend(auto_init=True):
    launch_ruby_svc("frontend-rb", "ruby/frontend", "rackup ./frontend.ru", auto_init=auto_init)

def launch_ruby_message_service(auto_init=True):
    launch_ruby_svc("message-rb", "ruby/message-service", "rackup message.ru --server puma", auto_init=auto_init)

def launch_ruby_name_service(auto_init=True):
    launch_ruby_svc("name-rb", "ruby/name-service", "ruby name.rb", auto_init=auto_init)

def launch_ruby_year_service(auto_init=True):
    launch_ruby_svc("year-rb", "ruby/year-service", "rackup --port 6001", auto_init=auto_init)

def launch_java_svc(name, dirname="", flags="", auto_init=True):
    '''
    Starts a single Java service.

    Parameters:
    name: used to display the name of the process in the tilt tab
    dirname: (optional) directory name in which to run `go run main.go` defaults to 'name'
    flags: (optional) any additional flags to add to the command line
    '''

    env = {
        'SERVICE_NAME': name,
        'OTEL_SERVICE_NAME': name
    }
    cmd = "cd {} && gradle bootRun".format(
        dirname if dirname else name,
        flags if flags else ""
    )
    print("About to start {} with command {}".format(name, cmd))
    local_resource(name, "", auto_init=auto_init, serve_cmd=cmd, serve_env=env)

def launch_java_frontend(auto_init=True):
    launch_java_svc("frontend-java", dirname="java/frontend", auto_init=auto_init)

def launch_java_message_service(auto_init=True):
    launch_java_svc("message-java", dirname="java/message-service", auto_init=auto_init)

def launch_java_name_service(auto_init=True):
    launch_java_svc("name-java", dirname="java/name-service", auto_init=auto_init)

def launch_java_year_service(auto_init=True):
    launch_java_svc("year-java", dirname="java/year-service", auto_init=auto_init)

def launch_dotnet_svc(name, dirname="", flags="", auto_init=True):
    '''
    Starts a single .NET service.

    Parameters:
    name: used to display the name of the process in the tilt tab
    dirname: (optional) directory name in which to run `go run main.go` defaults to 'name'
    flags: (optional) any additional flags to add to the command line
    '''

    cmd = "cd {} && dotnet run".format(
        dirname if dirname else name,
        flags if flags else ""
    )
    print("About to start {} with command {}".format(name, cmd))
    local_resource(name, "", auto_init=auto_init, serve_cmd=cmd)

def launch_dotnet_frontend(auto_init=True):
    launch_dotnet_svc("frontend-dotnet", dirname="dotnet/frontend", auto_init=auto_init)

def launch_dotnet_message_service(auto_init=True):
    launch_dotnet_svc("message-dotnet", dirname="dotnet/message-service", auto_init=auto_init)

def launch_dotnet_name_service(auto_init=True):
    launch_dotnet_svc("name-dotnet", dirname="dotnet/name-service", auto_init=auto_init)

def launch_dotnet_year_service(auto_init=True):
    launch_dotnet_svc("year-dotnet", dirname="dotnet/year-service", auto_init=auto_init)

def launch_node_svc(name, dirname="", flags="", auto_init=True):
    '''
    Starts a single Node service.

    Parameters:
    name: used to display the name of the process in the tilt tab
    dirname: (optional) directory name in which to run `npm start` defaults to 'name'
    flags: (optional) any additional flags to add to the command line
    '''
    
    env = {'SERVICE_NAME': name}

    cmd = "cd {} && npm install && npm start".format(
        dirname if dirname else name,
        flags if flags else ""
    )
    print("About to start {} with command {}".format(name, cmd))
    local_resource(name, "", auto_init=auto_init, serve_cmd=cmd, serve_env=env)

def launch_node_frontend(auto_init=True):
  launch_node_svc("frontend-node", dirname="node/frontend", auto_init=auto_init)

def launch_node_message_service(auto_init=True):
  launch_node_svc("message-node", dirname="node/message-service", auto_init=auto_init)

def launch_node_name_service(auto_init=True):
  launch_node_svc("name-node", dirname="node/name-service", auto_init=auto_init)

def launch_node_year_service(auto_init=True):
  launch_node_svc("year-node", dirname="node/year-service", auto_init=auto_init)


languages = ["go", "py", "rb", "java", "dotnet", "node"]

config.define_string_list("to-run", args=True)
cfg = config.parse()

groups = {
    'node': ['frontend-node', 'message-node', 'name-node', 'year-node'],
    'go': ['frontend-go', 'message-go', 'name-go', 'year-go']
}

resources = ['collector', 'redis', 'curl greeting']
to_run = cfg.get('to-run', []) or ["go"]

for arg in to_run:
  if arg in groups:
    resources += groups[arg]
  else:
    # also support specifying individual services instead of groups, e.g. `tilt up a b d`
    resources.append(arg)
config.set_enabled_resources(resources)

# can we run services programatically?






# Launch one of each of these types of services. Go services init by default
launch_go_frontend()
launch_python_frontend()
launch_ruby_frontend()
launch_java_frontend()
launch_dotnet_frontend()
launch_node_frontend()

launch_go_message_service()
launch_python_message_service()
launch_ruby_message_service()
launch_java_message_service()
launch_dotnet_message_service()
launch_node_message_service()

launch_go_name_service()
launch_python_name_service()
launch_ruby_name_service()
launch_java_name_service()
launch_dotnet_name_service()
launch_node_name_service()

launch_go_year_service()
launch_python_year_service()
launch_ruby_year_service()
launch_java_year_service()
launch_dotnet_year_service()
launch_node_year_service()

###
# Notes
###

# syntax for local_resource:
# local_resource ( name , build_cmd , deps=None , trigger_mode=TRIGGER_MODE_AUTO , resource_deps=[] , ignore=[] , auto_init=True , serve_cmd='go run cmd/shepherd/main.go -debug' )
# name
# command to build the thing to run (empty in our world)
# deps are a list of files to watch and, if changed, restart the process
# serve_cmd is the command to run to start the thing, and it's expected that it won't exit
# eg
# local_resource ( "shepherd" , "" , serve_cmd='cd cmd/shepherd && go run main.go -debug -p :8081' )

# link to tilt API docs: https://docs.tilt.dev/api.html
