FROM golang:1.18 AS build
WORKDIR /src
ENV CGO_ENABLED=0

COPY go.mod .

COPY frontend/. .
RUN go build -o /out/frontend .

FROM scratch AS bin
WORKDIR /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /out/frontend /app/
CMD ["/app/frontend"]
