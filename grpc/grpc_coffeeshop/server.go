package main

import (
	"context"
	pb "grpc/coffeeshop_proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCoffeeShopServer
}

func (s *server) GetMenu(menuRequest *pb.MenuRequest, srv pb.CoffeeShop_GetMenuServer) error {
	items := []*pb.Item{
		&pb.Item{
			Id:   "1",
			Name: "Black Coffee",
		},
		&pb.Item{
			Id:   "2",
			Name: "Americano",
		},
		&pb.Item{
			Id:   "3",
			Name: "Vanilla Latte",
		},
	}

	for i, _ := range items {
		srv.Send(&pb.Menu{
			Items: items[0 : i+1],
		})
	}

	return nil
}

func (s *server) PlaceOrder(ctx context.Context, order *pb.Order) (*pb.Receipt, error) {
	return &pb.Receipt{
		Id: "ABC123",
	}, nil
}

func (s *server) GetOrderStatus(ctx context.Context, receipt *pb.Receipt) (*pb.OrderStatus, error) {
	return &pb.OrderStatus{
		OrderId: receipt.Id,
		Status:  "IN PROGRESS",
	}, nil
}

func main() {
	// setup listener on port 9001
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("Failed to listen to port 9001: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterCoffeeShopServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9001: %v", err)
	}
}
