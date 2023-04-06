FROM golang:1.18 AS build
WORKDIR /src
ENV CGO_ENABLED=0

COPY go.mod .

COPY name-service/. .
RUN go build -o /out/name-service .

FROM scratch AS bin
WORKDIR /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /out/name-service /app/
CMD ["/app/name-service"]
