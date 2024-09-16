# stage 1
FROM golang:1.21-alpine3.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# stage 2
FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/main .
COPY .env .

EXPOSE 8000

CMD [ "./main" ]