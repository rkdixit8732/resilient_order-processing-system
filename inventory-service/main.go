package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/rkdixit8732/resilient_order-processing-system/proto/inventorypb"

	"google.golang.org/grpc"
)

type inventoryServer struct {
	pb.UnimplementedInventoryServiceServer
}

func (s *inventoryServer) ReserveItem(ctx context.Context, req *pb.InventoryRequest) (*pb.InventoryResponse, error) {
	log.Printf("Reserving item: %s, quantity: %d", req.ItemId, req.Quantity)
	return &pb.InventoryResponse{Success: true, Message: "Item reserved"}, nil
}

func (s *inventoryServer) ReleaseItem(ctx context.Context, req *pb.InventoryRequest) (*pb.InventoryResponse, error) {
	log.Printf("Releasing item: %s, quantity: %d", req.ItemId, req.Quantity)
	return &pb.InventoryResponse{Success: true, Message: "Item released"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterInventoryServiceServer(grpcServer, &inventoryServer{})
	fmt.Println("Inventory Service running on port 50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
