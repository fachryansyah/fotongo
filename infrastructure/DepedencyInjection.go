//go:build wireinject
// +build wireinject

package infrastructure

import (
	// "fotongo/app/utils/baseCommands"
	// "fotongo/infrastructure/services"
	prismaDB "fotongo/infrastructure/services/prisma/db"

	"github.com/google/wire"
	// "gorm.io/gorm"
)

func NewDatabaseProvider() (*prismaDB.PrismaClient, error) {
	wire.Build(prismaDB.NewClient)
	return &prismaDB.PrismaClient{}, nil
}
