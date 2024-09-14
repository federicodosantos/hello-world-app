# stage 1
FROM golang:1.21.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY ./cmd/main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin

# stage 2
FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/bin ./bin

EXPOSE 8000

CMD [ "/app/bin" ]