# Build stage
FROM golang:1.24 AS builder
WORKDIR /app

COPY ../proto ../proto
COPY . .

RUN go mod tidy
RUN go build -o payment-service .

# Run stage
FROM gcr.io/distroless/base-debian11
COPY --from=builder /app/payment-service /payment-service
ENTRYPOINT ["/payment-service"]