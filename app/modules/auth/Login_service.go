package auth

import (
	authDtos "fotongo/app/modules/auth/dtos"
	"fotongo/app/utils/baseCommands"
	"fotongo/app/utils/dtos"
	"fotongo/domains/entities"
	"log"

	"github.com/dgrijalva/jwt-go"
)

func (s *ServiceAuth) Login(request authDtos.LoginRequest) (dtos.JSONResponses, error) {
	jwtSecret := baseCommands.GetJWTConfig()

	var user entities.User

	tokenJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserID": user.ID,
		"Email":  user.Email,
		"admin":  false,
	})
	token, err := tokenJwt.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Println(err)
		return baseCommands.InternalServerErrorResponses(), err
	}

	return baseCommands.SuccessResponses(token), nil
}
