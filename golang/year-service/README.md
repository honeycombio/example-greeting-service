# kubernetifying year only

## setup

replace `<APIKEY>` with actual api key

```sh
export HONEYCOMB_API_KEY=<APIKEY>
kubectl create secret generic honeycomb --from-literal=api-key=$HONEYCOMB_API_KEY
```

```sh
# build the docker image for the year service
docker build -t year-go:local .

# deploy the service in k8s (sends straight to honeycomb)
kubectl apply -f year.yaml

# make sure everything is up and running
kubectl get pods

# follow logs
kubectl logs deployments/year-go --follow
```

`curl localhost:6001/year`
