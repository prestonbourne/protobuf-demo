package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/prestonbourne/protobuf-demo/protos/currency"
	"google.golang.org/grpc"
)

type Currency struct {
	log                                  *log.Logger
	currency.UnimplementedCurrencyServer //no idea what this does
}

func (c *Currency) GetRate(ctx context.Context, req *currency.RateRequest) (*currency.RateResponse, error) {
	c.log.Printf("%+v", req)

	return &currency.RateResponse{Rate: 0.5}, nil
}

func main() {
	// not sure what flag does here 🤔
	log := log.New(os.Stdout, "[grpc-demo]", 0)

	gs := grpc.NewServer()
	cs := &Currency{log, currency.UnimplementedCurrencyServer{}}

	//GetRate(context.Context, *RateRequest) (*RateResponse, error)
	currency.RegisterCurrencyServer(gs, cs)

	listener, err := net.Listen("tcp", ":3001")
	if err != nil {
		fmt.Println(err.Error())
	}

	gs.Serve(listener)
}
