# auto-instrument go services

## temporary start

make local docker image for auto-instrumentation agent:

```sh
# clone repo
git clone git@github.com:open-telemetry/opentelemetry-go-instrumentation.git
# navigate into new repo
cd opentelemetry-go-instrumentation
# make docker image called otel-go-agent:v0.1
make docker-build IMG=otel-go-agent:v0.1
# make sure you have it locally
docker images | grep otel-go-agent
```

## setup

```sh
# replace `<APIKEY>` with actual api key
export HONEYCOMB_API_KEY=<APIKEY>

# build the docker images for all the services
docker-compose build

# deploy the greetings namespace and services in k8s
kubectl apply -k greetings/

# create secret with api key
kubectl create secret generic honeycomb --from-literal=api-key=$HONEYCOMB_API_KEY -n greetings

# deploy the services with the auto-instrumentation agent
kubectl apply -f greetings-instrumented.yaml -n greetings

# deploy the collector
kubectl apply -f otel-collector.yaml -n greetings

# make sure everything is up and running
kubectl get pods -n greetings

# follow logs for collector (optional)
kubectl logs deployments/otel-collector --follow -n greetings
```

`curl localhost:7007/greeting`

## cleanup

```sh
# delete secret with api key
kubectl delete namespace greetings

# make sure everything is gone
kubectl get pods -n greetings
```
