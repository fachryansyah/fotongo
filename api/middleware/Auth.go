package middleware

import (
	"fmt"
	"fotongo/app/utils/baseCommands"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func ProtectedUser(c fiber.Ctx) error {
	tokenString := c.Get("Authentication")
	jwtSecret := baseCommands.GetJWTConfig()

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(baseCommands.InternalServerErrorResponses())
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusForbidden).JSON(baseCommands.InternalServerErrorResponses())
	}

	userID := claims["identity"].(map[string]interface{})["ID"]
	email := claims["identity"].(map[string]interface{})["Email"]
	name := claims["identity"].(map[string]interface{})["Name"]

	c.Set("UserID", fmt.Sprint(userID))
	c.Set("Email", fmt.Sprint(email))
	c.Set("Name", fmt.Sprint(name))

	return c.Next()
}
