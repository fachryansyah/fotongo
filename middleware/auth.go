package middleware

import (
	"fmt"
	"fotongo/utils"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func ProtectedUser(c *fiber.Ctx) error {
	jwtKey := os.Getenv("JWT_KEY")
	tokenString := c.Get("Authorization")

	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtKey), nil
	})
	if err != nil {
		log.Println(err)
		return utils.ResponseUnauthenticated(c, fmt.Sprint(err), "Error")
	}
	return c.Next()
}
