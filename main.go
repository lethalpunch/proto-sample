package main

import (
	"log"
	"net"
	"time"

	"github.com/maximilienandile/demo-grpc/invoicer"
	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

//	func (s myInvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
//		return &invoicer.CreateResponse{
//			Currency: req.From,
//			Sender:   req.To,
//		}, nil
//	}
func (s myInvoicerServer) Create(req *invoicer.CreateRequest, server invoicer.Invoicer_CreateServer) error {
	for i := 0; i < 10; i++ {
		server.Send(&invoicer.CreateResponse{
			Currency: req.From,
			Sender:   req.To,
		})
		time.Sleep(2 * time.Second)
	}
	return nil

}
func main() {
	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}
	serverRegistrar := grpc.NewServer()
	service := &myInvoicerServer{}
	invoicer.RegisterInvoicerServer(serverRegistrar, service)
	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}
