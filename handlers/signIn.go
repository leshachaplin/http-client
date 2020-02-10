package handlers

import (
	"context"
	"github.com/labstack/echo"
	"github.com/leshachaplin/grpc-server/protocol"
	"log"
	"net/http"
)

func (a *Auth) SignIn(c echo.Context) error {
	user := new(UserModel)

	if err := c.Bind(user); err != nil { //The default binder supports decoding application/json,
		// application/xml and application/x-www-form-urlencoded data based on the Content-Type header.
		return err
	}
	requestAuth := &protocol.SignInRequest{Login: user.Login, Password: user.Password}
	responseAuth, err := a.client.SignIn(context.Background(), requestAuth)
	if err != nil {
		log.Println(err)
		return echo.ErrUnauthorized
	}
	tokenString := responseAuth.GetToken()

	return c.JSON(http.StatusOK, &TokenModel{Token: tokenString})
}
