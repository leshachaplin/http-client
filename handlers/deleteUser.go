package handlers

import (
	"context"
	"github.com/labstack/echo"
	"github.com/leshachaplin/grpc-server/protocol"
	"log"
	"net/http"
)

func (a *Auth) DeleteUser(c echo.Context) error {
	user := new(UserModel)
	if err := c.Bind(user); err != nil { //The default binder supports decoding application/json,
		// application/xml and application/x-www-form-urlencoded data based on the Content-Type header.
		return err
	}
	// I should check there, if the current user has any rights to delete this user
	claims := c.Get("claims").(map[string]string)

	if claims["admin"] == "true" {
		requestToDelete := &protocol.DeleteRequest{Login: user.Login}
		_, err := a.client.Delete(context.Background(), requestToDelete)
		// If there are some problems, so user has made a mistake
		if err != nil {
			log.Println(err)
			return echo.ErrBadRequest

		}
	} else {
		return echo.ErrUnauthorized
	}
	return c.String(http.StatusOK, "")
}
