# to be built from docker-compose
FROM golang:1.24 AS build
WORKDIR /src
ENV CGO_ENABLED=0

COPY go.mod .
COPY frontend/. .
RUN go get
RUN go build -o /out/frontend .

FROM scratch AS bin
WORKDIR /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /out/frontend /app/

# local code uses localhost:9000 and localhost:8000
# image will use the below endpoints
ENV MESSAGE_ENDPOINT=http://message:9000
ENV NAME_ENDPOINT=http://name:8000
# or override at runtime in k8s or docker-compose

CMD ["/app/frontend"]
