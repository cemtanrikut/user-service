# Builder image
FROM golang:1.18-alpine AS builder

# Load dependencies and tools
RUN apk --no-cache add git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /user-service ./cmd/user-service

FROM alpine:latest

WORKDIR /app

COPY --from=builder /user-service .

CMD ["./user-service"]

EXPOSE 8080
