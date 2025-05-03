FROM golang:1.24 AS builder
WORKDIR /app

COPY ../go.mod ../go.sum ./
RUN go mod download

COPY ../ ./
WORKDIR /app/inventory-service
RUN go build -o inventory-service .

FROM gcr.io/distroless/base-debian11

WORKDIR /app
COPY --from=builder /app/inventory-service/inventory-service .
EXPOSE 50052
ENTRYPOINT ["/app/inventory-service"]