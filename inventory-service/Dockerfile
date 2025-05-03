# -------- Build Stage --------
    FROM golang:1.24 AS builder

    WORKDIR /app
    
    # Install proto compiler plugins
    ENV GO_BIN_PATH=/go/bin
    ENV PATH=${GO_BIN_PATH}:${PATH}
    ENV GOPROXY=https://proxy.golang.org,direct
    
    RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@vlatest
    
    RUN export PATH="$PATH:$(go env GOPATH)/bin"

    # Copy go module files
    COPY go.mod go.sum ./
    RUN go mod download
    
    # Copy source and proto files
    COPY . .
    #COPY proto ./proto
    
    # Generate .pb.go files
    RUN protoc --go_out=. --go-grpc_out=. proto/inventory.proto

    # Build the binary
    RUN go build -o inventory-service .
    
    # -------- Runtime Stage --------
    FROM gcr.io/distroless/base-debian11
    COPY --from=builder /app/inventory-service /inventory-service
    
    ENTRYPOINT ["/inventory-service"]