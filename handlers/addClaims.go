package handlers

import (
	"context"
	"github.com/labstack/echo"
	"github.com/leshachaplin/grpc-server/protocol"
	"net/http"
)

func (a *Auth) AddClaims(c echo.Context) error {
	claim := new(ClaimModel)
	claims := c.Get("claims").(map[string]string)
	if err := c.Bind(claim); err != nil {
		return err
	}
	claims[claim.Key] = claim.Value
	requestToAddClaim := &protocol.AddClaimsRequest{Claims: claims}
	_, err := a.client.AddClaims(context.Background(), requestToAddClaim)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "")
}
