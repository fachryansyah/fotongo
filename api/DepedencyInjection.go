//go:generate go run github.com/google/wire/cmd/wire
//go:build wireinject
// +build wireinject

package api

import (
	"fotongo/api/controllers"
	"fotongo/app/modules/auth"
	"fotongo/infrastructure"

	"github.com/google/wire"
)

func InitializeAuthController() (*controllers.AuthController, error) {
	wire.Build(infrastructure.NewDatabaseProvider, auth.NewAuthService, controllers.NewAuthController)
	return &controllers.AuthController{}, nil
}
