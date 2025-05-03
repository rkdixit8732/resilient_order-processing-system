FROM golang:1.24 AS builder
WORKDIR /app

COPY ../go.mod ../go.sum ./
RUN go mod download

COPY ../ ./
WORKDIR /app/payment-service
RUN go build -o payment-service .

FROM gcr.io/distroless/base-debian11

WORKDIR /app
COPY --from=builder /app/payment-service/payment-service .
EXPOSE 50051
ENTRYPOINT ["/app/payment-service"]