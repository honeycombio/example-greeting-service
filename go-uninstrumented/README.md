# example greeting service in golang - uninstrumented

from root directory:

```sh
# build docker images
docker-compose build
# TODO publish docker image to skip docker-compose step
# deploy the services in k8s
kubectl apply -f greetings.yaml
# port-forward the frontend service
kubectl port-forward svc/frontend 7777:7777
```

Then in a new terminal, `curl localhost:7777/greeting`

When all done:

```sh
kubectl delete -f greetings.yaml
```
