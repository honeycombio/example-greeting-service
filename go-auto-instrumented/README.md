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

replace `<APIKEY>` with actual api key

```sh
export HONEYCOMB_API_KEY=<APIKEY>
```

```sh
# build the docker images for all the services
docker-compose build

# create secret with api key
kubectl create secret generic honeycomb --from-literal=api-key=$HONEYCOMB_API_KEY

# deploy the service in k8s
kubectl apply -f greetings.yaml

# deploy the services with the auto-instrumentation agent
kubectl apply -f greetings-instrumented.yaml

# deploy the collector
kubectl apply -f otel-collector.yaml

# make sure everything is up and running
kubectl get pods

# follow logs for collector (optional)
kubectl logs deployments/otel-collector --follow
```

`curl localhost:7007/greeting`

## cleanup

```sh
# delete secret with api key
kubectl delete secret honeycomb

# delete the services with the auto-instrumentation agent
kubectl delete -f greetings-instrumented.yaml

# delete the collector service
kubectl delete -f otel-collector.yaml

# make sure everything is gone
kubectl get pods
```
