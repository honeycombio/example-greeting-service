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
ENV MESSAGE_ENDPOINT=http://message-go-svc:9000
ENV NAME_ENDPOINT=http://name-go-svc:8000
ENV YEAR_ENDPOINT=http://year-go-svc:6001
CMD ["/app/message-service"]
