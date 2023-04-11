# auto-instrument go services

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
