package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/rkdixit8732/resilient_order-processing-system/proto/paymentpb"

	"google.golang.org/grpc"
)

type paymentServer struct {
	pb.UnimplementedPaymentServiceServer
}

func (s *paymentServer) ProcessPayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	log.Printf("Processing payment for order: %s", req.OrderId)
	return &pb.PaymentResponse{Success: true, Message: "Payment processed"}, nil
}

func (s *paymentServer) RefundPayment(ctx context.Context, req *pb.RefundRequest) (*pb.PaymentResponse, error) {
	log.Printf("Refunding payment for order: %s", req.OrderId)
	return &pb.PaymentResponse{Success: true, Message: "Payment refunded"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPaymentServiceServer(grpcServer, &paymentServer{})
	fmt.Println("Payment Service running on port 50053")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
