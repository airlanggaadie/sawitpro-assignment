package handler

import (
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
)

// This is just a health check endpoint to get service information.
// (GET /healthz)
func (s *Server) Healthz(ctx echo.Context) error {
	response := generated.HealthCheckResponse{Status: "OK"}
	return ctx.JSON(http.StatusOK, response)
}
