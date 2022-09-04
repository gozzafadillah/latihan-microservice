package middlewares

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustomClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

type ConfigJwt struct {
	SecretJWT string
}

func (jwtConf *ConfigJwt) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
		ErrorHandlerWithContext: middleware.JWTErrorHandlerWithContext(func(e error, c echo.Context) error {
			return c.JSON(401, map[string]interface{}{
				"message": "unauthorized, missing or malformed jwt",
				"rescode": 401,
			})
		}),
	}
}

// GenerateToken jwt ...
func (jwtConf *ConfigJwt) GenerateToken(userID string, email string, role string) (string, error) {
	claims := JwtCustomClaims{
		ID:    userID,
		Email: email,
		Role:  role,
	}

	// Create token with claims
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(jwtConf.SecretJWT))

	return token, err
}

// GetUser from jwt ...
func GetUser(c echo.Context) *JwtCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return claims
}
