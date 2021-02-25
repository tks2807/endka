package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	pb "endka/proto"
)


type Server struct {
}

func (s *Server) PrimeNumberDecomposition(ctx context.Context, req *pb.PrimeRequest) (*pb.PrimeReply, error) {
	input := in.N1
	var ans []int
	for ;input%2==0;input=input/2 {
		fmt.Print(2, " ")
		ans = append(ans, 2)
	}
	i := 3
	for ;i<=int(math.Sqrt(float64(n))); i = i + 2 {
		for ;input%i==0; {
			fmt.Print(i, " ")
			ans = append(ans, i)
			input = input/i
		}
	}
	&pb.PrimeReply{N1: ans}, nil
}

func (s *Server) ComputeAverage(ctx context.Context, req *pb.AverageRequest) (*pb.AverageReply, error) {
	return &pb.AverageReply{N1: }, nil
}




func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50050")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})
	log.Println("Server is running on port:50050")
	if err := s.Server(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}

