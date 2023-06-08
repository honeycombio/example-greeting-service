# example greeting service in golang - uninstrumented

Images are now available on [GitHub Packages](https://github.com/orgs/honeycombio/packages?repo_name=example-greeting-service)

## Steps to Build Locally

```sh
# build docker images
docker-compose build
```

## Steps to Run in Kubernetes

```sh
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
