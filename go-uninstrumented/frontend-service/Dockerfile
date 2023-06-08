# to be built from docker-compose or publish-ghcr workflow
FROM golang:1.20 AS build
WORKDIR /src
ENV CGO_ENABLED=0

COPY go.mod .
COPY frontend-service/. .
RUN go get
RUN go build -o /out/frontend-service .

FROM scratch AS bin
WORKDIR /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /out/frontend-service /app/

# local code uses localhost:9000 and localhost:8000
# image will use the below endpoints
ENV MESSAGE_ENDPOINT=http://message:9000
ENV NAME_ENDPOINT=http://name:8000
# or override at runtime in k8s or docker-compose

CMD ["/app/frontend-service"]
