package provider

import (
	"fmt"
	"fotongo/model"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) (*model.User, error) {
	jwtKey := os.Getenv("JWT_KEY")
	tokenString := c.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtKey), nil
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	userID := claims["identity"].(map[string]interface{})["id"]
	email := claims["identity"].(map[string]interface{})["email"]
	name := claims["identity"].(map[string]interface{})["name"]

	return &model.User{
		ID:    fmt.Sprint(userID),
		Name:  fmt.Sprint(name),
		Email: fmt.Sprint(email),
	}, nil
}
