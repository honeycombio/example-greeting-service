# example greeting service in golang - uninstrumented

Images are now available on [GitHub Packages](https://github.com/orgs/honeycombio/packages?repo_name=example-greeting-service)

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

## Steps to Build Locally

To use local images, build using `docker-compose` and update the `greetings.yaml` with the new image names.

```sh
# build docker images
docker-compose build
```

As an example:

```yaml
    spec:
      serviceAccountName: year-go
      terminationGracePeriodSeconds: 0
      containers:
        - name: year
          imagePullPolicy: IfNotPresent
          image: egs-year-go:local
          ports:
          - containerPort: 6001
            name: http
```
