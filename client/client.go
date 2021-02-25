package client


import (
	"context"
	"fmt"
	"io"
	"log"
	"time"
	pb "endka/proto"
	"google.golang.org/grpc"
)



func main() {
	fmt.Println("Hello I'm a client")

	conn, err := grpc.Dial("localhost:50050", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)
	doPrimeRequest(c)
	doAverageRequest(c)
}

