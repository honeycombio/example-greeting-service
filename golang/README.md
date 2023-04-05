# kubernetifying go services

## setup

replace `<APIKEY>` with actual api key

```sh
export HONEYCOMB_API_KEY=<APIKEY>
kubectl create secret generic honeycomb --from-literal=api-key=$HONEYCOMB_API_KEY
```

```sh
# build the docker images
docker-compose build

# deploy the services in k8s
kubectl apply -f greetings.yaml
# deploy collector in k8s
kubectl apply -f otel-collector.yaml

# port-forward services so they are accessible
kubectl port-forward svc/frontend-golang 7007:7007
kubectl port-forward svc/message-golang 9000:9000
kubectl port-forward svc/name-golang 8000:8000
kubectl port-forward svc/year-golang 6001:6001

# make sure everything is up and running
kubectl get pods
```

`curl localhost:7007/greeting`

## cleanup

```sh
kubectl delete -f greetings.yaml
kubectl delete -f otel-collector.yaml
```
