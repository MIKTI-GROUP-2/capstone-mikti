// middlewares.go

package middlewares

import (
	"capstone-mikti/configs"
	"capstone-mikti/helper/jwt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Middleware struct {
	jwtService jwt.JWTInterface
	config     *configs.ProgrammingConfig
}

func NewMiddleware(config *configs.ProgrammingConfig) *Middleware {
	jwtService := jwt.NewJWT(config)

	return &Middleware{
		jwtService: jwtService,
		config:     config,
	}
}

func (m *Middleware) RoleMiddleware(allowedRole string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims, err := m.jwtService.ExtractToken(c)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
			}

			// Periksa role
			if claims.Role != allowedRole {
				return c.JSON(http.StatusForbidden, map[string]string{"error": "Akses ditolak, Masalah pada role"})
			}

			c.Set("user", claims)

			return next(c)
		}
	}
}
