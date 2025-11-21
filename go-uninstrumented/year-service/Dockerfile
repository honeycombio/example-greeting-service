# to be built from docker-compose or publish-ghcr workflow
FROM golang:1.24 AS build
WORKDIR /src
ENV CGO_ENABLED=0

COPY go.mod .

COPY year-service/*.go .
RUN go build -o /out/year-service .

FROM scratch AS bin
WORKDIR /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /out/year-service /app/
CMD ["/app/year-service"]
