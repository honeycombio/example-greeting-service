FROM golang:1.16-alpine AS build
WORKDIR /src
ENV CGO_ENABLED=0
COPY . .
RUN go build -o /out/message-service .

FROM alpine AS bin
WORKDIR /app
COPY --from=build /out/message-service /app/
CMD ["/app/message-service"]
