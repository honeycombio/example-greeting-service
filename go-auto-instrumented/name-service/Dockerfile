FROM golang:1.17 AS build
WORKDIR /src
ENV CGO_ENABLED=0

COPY go.mod .

COPY name-service/. .
RUN go build -o /out/name-service .

FROM scratch AS bin
WORKDIR /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /out/name-service /app/
ENV MESSAGE_ENDPOINT=http://message-go-svc:9000
ENV NAME_ENDPOINT=http://name-go-svc:8000
ENV YEAR_ENDPOINT=http://year-go-svc:6001
CMD ["/app/name-service"]
