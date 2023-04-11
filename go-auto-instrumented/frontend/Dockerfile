FROM golang:1.20 AS build
WORKDIR /src
ENV CGO_ENABLED=0

COPY go.mod .
COPY frontend/. .
RUN go get
RUN go build -o /out/frontend .

FROM scratch AS bin
WORKDIR /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /out/frontend /app/
ENV MESSAGE_ENDPOINT=http://message-go-svc:9000
ENV NAME_ENDPOINT=http://name-go-svc:8000
ENV YEAR_ENDPOINT=http://year-go-svc:6001
CMD ["/app/frontend"]
