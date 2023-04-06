FROM golang:1.18 AS build
WORKDIR /src
ENV CGO_ENABLED=0

COPY go.mod .

COPY message-service/. .
RUN go build -o /out/message-service .

FROM scratch AS bin
WORKDIR /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /out/message-service /app/
CMD ["/app/message-service"]
