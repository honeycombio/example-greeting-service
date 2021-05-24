FROM golang:1.16-alpine AS build
WORKDIR /src
ENV CGO_ENABLED=0
COPY . .
RUN go build -o /out/name-service ./main.go

FROM scratch AS bin
WORKDIR /app
COPY --from=build /out/name-service /app/
CMD ["/app/name-service"]
