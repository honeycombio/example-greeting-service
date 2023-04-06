# kubernetifying year only

## setup

replace `<APIKEY>` with actual api key

```sh
export HONEYCOMB_API_KEY=<APIKEY>
```

```sh
# build the docker image for the year service
docker build -t year-go-auto:local .

# create secret with api key
kubectl create secret generic honeycomb --from-literal=api-key=$HONEYCOMB_API_KEY

# deploy the service in k8s
kubectl apply -f year.yaml

# deploy the service with the auto-instrumentation agent
kubectl apply -f year-instrumented.yaml

# deploy the collector
kubectl apply -f otel-collector.yaml

# make sure everything is up and running
kubectl get pods

# follow logs for collector (optional)
kubectl logs deployments/otel-collector --follow
```

`curl localhost:6001/year`

## cleanup

```sh
# delete secret with api key
kubectl delete secret honeycomb

# delete the service in k8s
kubectl delete -f year.yaml

# delete the service with the auto-instrumentation agent
kubectl delete -f year-instrumented.yaml

# make sure everything is gone
kubectl get pods
```
