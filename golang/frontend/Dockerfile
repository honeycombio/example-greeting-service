FROM golang:1.20 AS build
WORKDIR /src
ENV CGO_ENABLED=0

COPY go.mod .
COPY go.sum .
COPY frontend/. .

RUN go get
RUN go build -o /out/frontend .

FROM scratch AS bin
WORKDIR /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /out/frontend /app/
CMD ["/app/frontend"]
