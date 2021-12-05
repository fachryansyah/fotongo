package auth

import (
	"fotongo/infrastructure/services/prisma/db"
)

type ServiceAuth struct {
	db *db.PrismaClient
}

func NewAuthService(db *db.PrismaClient) AuthInterface {
	return &ServiceAuth{
		db,
	}
}
