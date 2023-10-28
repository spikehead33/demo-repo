FROM golang:1.21.3-bullseye AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/server

FROM debian:buster

WORKDIR /

COPY --from=builder /app/bin/server /bin/server
EXPOSE 3000
ENTRYPOINT ["/bin/server"]
