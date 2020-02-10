package handlers

import (
	"context"
	"github.com/labstack/echo"
	"github.com/leshachaplin/grpc-server/protocol"
	"net/http"
)

func (a *Auth) DeleteClaims(c echo.Context) error {
	claim := new(ClaimModel)
	claims := c.Get("claims").(map[string]string)
	if err := c.Bind(claim); err != nil {
		return err
	}
	delete(claims, claim.Key)
	requestToDeleteClaim := &protocol.DeleteClaimsRequest{Claims: claims}
	_, err := a.client.DeleteClaims(context.Background(), requestToDeleteClaim)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "")
}
