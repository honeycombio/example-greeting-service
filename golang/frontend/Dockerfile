FROM golang:1.16-alpine AS build
WORKDIR /src
ENV CGO_ENABLED=0
COPY . .
RUN go get
RUN go build -o /out/frontend .

FROM scratch AS bin
WORKDIR /app
COPY --from=build /out/frontend /app/
CMD ["/app/frontend"]
