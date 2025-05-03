FROM golang:1.24 AS builder
WORKDIR /app

COPY ../go.mod ../go.sum ./
RUN go mod download

COPY ../ ./
WORKDIR /app/order-service
RUN go build -o order-service .

FROM gcr.io/distroless/base-debian11

WORKDIR /app
COPY --from=builder /app/order-service/order-service .
EXPOSE 50050

ENTRYPOINT ["/app/order-service"]