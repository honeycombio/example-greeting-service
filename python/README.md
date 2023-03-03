# Python Greeting Services

The four different services in example-greeting-service have been implemented here in Python, each of them using a different Python web framework.

pyenv and poetry are being used to manage dependencies.

## How to run

### Option 1: All Python Services via Tilt

Prerequisite: `brew install poetry` to run locally if not already installed.

In the top level directory run `tilt up py`

### Option 2: All Python Services via Docker-Compose

In the python directory run `docker-compose up --build`

### Option 3:  One Python Service at a time via Docker

In each service directory, run the Docker commands that build and start the service with the corresponding service names and ports.

` docker build -t frontend .`
` docker run -dp 7007:7007 frontend`

Or however it is you prefer to Docker. :)
