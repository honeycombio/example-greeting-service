FROM golang:1.17-alpine AS build
WORKDIR /src
ENV CGO_ENABLED=0
COPY . .
RUN go build -o /out/year-service ./main.go

FROM scratch AS bin
WORKDIR /app
COPY --from=build /out/year-service /app/
CMD ["/app/year-service"]
