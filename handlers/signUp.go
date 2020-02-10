package handlers

import (
	"context"
	"github.com/labstack/echo"
	"github.com/leshachaplin/grpc-server/protocol"
	"net/http"
)

func (a *Auth) SignUp(c echo.Context) error {
	user := new(UserModel)
	if err := c.Bind(user); err != nil {
		return err
	}
	requestRegistration := &protocol.SignInRequest{Login: user.Login, Password: user.Password}
	responseRegistration, err := a.client.SignIn(context.Background(), requestRegistration)
	if err != nil {
		return err
	}
	tokenString := responseRegistration.GetToken()

	return c.JSON(http.StatusOK, &TokenModel{Token: tokenString})
}
