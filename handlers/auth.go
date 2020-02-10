package handlers

import "github.com/leshachaplin/grpc-server/protocol"

type Auth struct {
	client protocol.AuthServiceClient
}

func NewHandler(client protocol.AuthServiceClient) *Auth {
	return &Auth{client: client}
}
