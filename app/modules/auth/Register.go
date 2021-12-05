package auth

import (
	"context"
	authDtos "fotongo/app/modules/auth/dtos"
	"fotongo/app/utils/baseCommands"
	"fotongo/app/utils/dtos"
	"fotongo/infrastructure/services/prisma/db"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (s *ServiceAuth) Register(request authDtos.RegisterRequest) (dtos.JSONResponses, error) {

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	if err != nil {
		return baseCommands.InternalServerErrorResponses(), err
	}

	log.Println(string(hashPassword))

	ctx := context.Background()
	user, err := s.db.User.CreateOne(
		db.User.Username.Set(request.Username),
		db.User.Email.Set(request.Email),
		db.User.Password.Set(string(hashPassword)),
		db.User.Roles.Set("USER"),
	).Exec(ctx)
	if err != nil {
		log.Println("error cuy")
		log.Println(err)
		return baseCommands.InternalServerErrorResponses(), err
	}

	return baseCommands.SuccessResponses(authDtos.RegisterResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}), nil
}
