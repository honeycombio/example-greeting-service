# to be built from docker-compose
FROM golang:1.20 AS build
WORKDIR /src
ENV CGO_ENABLED=0

COPY go.mod .

COPY name-service/. .
RUN go build -o /out/name-service .

FROM scratch AS bin
WORKDIR /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /out/name-service /app/

# local code uses localhost:6001/year
# image will use the below endpoint
ENV YEAR_ENDPOINT=http://year:6001
# or override at runtime in k8s or docker-compose

CMD ["/app/name-service"]
