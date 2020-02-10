package main

import (
	"github.com/ivansukach/http-server/config"
	"github.com/ivansukach/http-server/handlers"
	"github.com/ivansukach/http-server/middlewares"
	"github.com/labstack/echo"
	"github.com/leshachaplin/grpc-server/protocol"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Клиент запущен ...")
	opts := grpc.WithInsecure() //WithInsecure returns a DialOption which disables transport security for this ClientConn.
	// Note that transport security is required unless WithInsecure is set.
	clientConnInterface, err := grpc.Dial(config.Localhost, opts) //attempt to connect to grpc-server
	if err != nil {
		log.Error(err)
	}
	defer clientConnInterface.Close() //A defer statement defers the execution of a function until the surrounding function returns.
	client := protocol.NewAuthServiceClient(clientConnInterface)
	auth := handlers.NewHandler(client)
	jwt := middlewares.NewJWT(client)
	e := echo.New()
	e.POST("/signIn", auth.SignIn)
	e.POST("/signUp", auth.SignUp)
	e.POST("/delete", auth.DeleteUser, jwt.Middleware)
	e.POST("/addClaims", auth.AddClaims, jwt.Middleware)
	e.POST("/deleteClaims", auth.DeleteClaims, jwt.Middleware)
	e.Logger.Fatal(e.Start(":8081"))
}
