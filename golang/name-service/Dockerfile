FROM golang:1.24 AS build
WORKDIR /src
ENV CGO_ENABLED=0

COPY go.mod .
COPY go.sum .
COPY name-service/. .

RUN go build -o /out/name-service ./main.go

FROM scratch AS bin
WORKDIR /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /out/name-service /app/
CMD ["/app/name-service"]
