FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app2

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/app2 .

EXPOSE 8081

CMD ["./app2"]