FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app1

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/app1 .

EXPOSE 8080

CMD ["./app1"]