package main

import (
	"context"
	"log"
	"net"

	pb "resilient_order-processing-system/proto/inventory"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedInventoryServiceServer
}

func (s *server) ReserveItem(ctx context.Context, req *pb.InventoryRequest) (*pb.InventoryResponse, error) {
	log.Printf("Reserving item: %v", req.ItemId)
	return &pb.InventoryResponse{
		Status:  "RESERVED",
		Message: "Item reserved successfully",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterInventoryServiceServer(s, &server{})
	log.Println("InventoryService listening on :50053")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}