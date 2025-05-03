package main

import (
	"context"
	"fmt"
	"log"
	"net"

	inventorypb "github.com/rkdixit8732/resilient_order-processing-systemproto/inventorypb"
	orderpb "rgithub.com/rkdixit8732/resilient_order-processing-system/proto/orderpb"
	paymentpb "github.com/rkdixit8732/resilient_order-processing-system/proto/paymentpb"

	"google.golang.org/grpc"
)

type orderServer struct {
	orderpb.UnimplementedOrderServiceServer
	inventoryClient inventorypb.InventoryServiceClient
	paymentClient   paymentpb.PaymentServiceClient
}

func (s *orderServer) CreateOrder(ctx context.Context, req *orderpb.OrderRequest) (*orderpb.OrderResponse, error) {
	log.Println("Creating order")

	_, err := s.inventoryClient.ReserveItem(ctx, &inventorypb.InventoryRequest{
		ItemId:   req.ItemId,
		Quantity: req.Quantity,
	})
	if err != nil {
		return &orderpb.OrderResponse{Status: "Inventory failed"}, err
	}

	_, err = s.paymentClient.ProcessPayment(ctx, &paymentpb.PaymentRequest{
		OrderId: req.UserId + "-order",
		UserId:  req.UserId,
		Amount:  100.0,
	})
	if err != nil {
		s.inventoryClient.ReleaseItem(ctx, &inventorypb.InventoryRequest{
			ItemId:   req.ItemId,
			Quantity: req.Quantity,
		})
		return &orderpb.OrderResponse{Status: "Payment failed"}, err
	}

	return &orderpb.OrderResponse{OrderId: "order123", Status: "Success"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	inventoryConn, _ := grpc.Dial("inventory-service:50052", grpc.WithInsecure())
	paymentConn, _ := grpc.Dial("payment-service:50053", grpc.WithInsecure())

	orderpb.RegisterOrderServiceServer(grpcServer, &orderServer{
		inventoryClient: inventorypb.NewInventoryServiceClient(inventoryConn),
		paymentClient:   paymentpb.NewPaymentServiceClient(paymentConn),
	})

	fmt.Println("Order Service running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
