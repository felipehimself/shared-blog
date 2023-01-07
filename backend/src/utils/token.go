package utils

import (
	"errors"
	"strconv"

	"shared-blog-backend/src/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func IsTokenValid(c *fiber.Ctx) (uint64, error) {

	cookie := c.Cookies("shared-blog-jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(*jwt.Token) (interface{}, error) {
		return []byte(config.SECRET_KEY), nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims := token.Claims.(*jwt.StandardClaims)
	issuer, err := strconv.ParseUint(claims.Issuer, 10, 64)

	if err != nil {
		return 0, err
	}

	return issuer, nil

}
