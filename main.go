package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/leshachaplin/http-client/protocol"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello client ...")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts, grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := protocol.NewHelloServiceClient(cc)
	request := &protocol.HelloRequest{Name: "lesha"}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		resp, _ := client.Hello(context.Background(), request)
		return c.String(http.StatusOK, resp.Greeting)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
