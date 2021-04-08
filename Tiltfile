# prints show up in the Tiltfile log so you know what's been run
print("hello my friends")


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
    print("About to start {} with command {}".format(name, cmd))
    local_resource(name, "", auto_init=auto_init, serve_cmd=cmd)


def launch_go_frontend(auto_init=True):
    launch_go_svc("go-frontend", dirname="golang/frontend")


def launch_go_message_service(auto_init=True):
    launch_go_svc("go-message", dirname="golang/message-service")


def launch_go_name_service(auto_init=True):
    launch_go_svc("go-name", dirname="golang/name-service")


def launch_go_year_service(auto_init=True):
    launch_go_svc("go-year", dirname="golang/year-service")


def launch_python_frontend(auto_init=True):
    cmd = "cd python/frontend && poetry install --no-root && poetry run python -m frontend"
    local_resource("py-frontend", "", auto_init=auto_init, serve_cmd=cmd)


def launch_python_message_service(auto_init=True):
    cmd = "cd python/message-service && poetry install --no-root && poetry run python -m message_service"
    local_resource("py-message", "", auto_init=auto_init, serve_cmd=cmd)


def launch_python_name_service(auto_init=True):
    cmd = "cd python/name-service && poetry install --no-root && poetry run flask run"
    local_resource("py-name", "", auto_init=auto_init, serve_cmd=cmd)


def launch_python_year_service(auto_init=True):
    cmd = "cd python/year-service && poetry install --no-root && poetry run yearservice/manage.py runserver 127.0.0.1:6001"
    local_resource("py-year", "", auto_init=auto_init, serve_cmd=cmd)

def launch_ruby_frontend(auto_init=True):
    cmd = "cd ruby/frontend && rackup ./frontend.ru"
    local_resource("rb-frontend", "", auto_init=auto_init, serve_cmd=cmd)

def launch_ruby_name_service(auto_init=True):
    cmd = "cd ruby/name-service && ruby name.rb"
    local_resource("rb-name", "", auto_init=auto_init, serve_cmd=cmd)


# Launch one of each of these types of services. Go services init by default
launch_go_frontend()
# launch_python_frontend()
# launch_ruby_frontend()

launch_go_message_service()
# launch_python_message_service()

launch_go_name_service()
# launch_python_name_service()
# launch_ruby_name_service()

launch_go_year_service()
# launch_python_year_service()

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

# link to quip doc on tilt: https://honeycomb.quip.com/h2MFAEUaKTKe
# link to tilt API docs: https://docs.tilt.dev/api.html
