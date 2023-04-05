FROM golang:1.20 AS build
WORKDIR /src
ENV CGO_ENABLED=0
COPY . .
RUN go get
RUN go build -o /out/year-service

FROM scratch AS bin
WORKDIR /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /out/year-service /app/
CMD ["/app/year-service"]
