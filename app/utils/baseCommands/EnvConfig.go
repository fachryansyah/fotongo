package baseCommands

import (
	"fotongo/app/utils/constant"
	"os"

	"gorm.io/driver/postgres"
)

func GetPostgreConfig() postgres.Config {
	configPostgre := postgres.Config{
		PreferSimpleProtocol: false,
		DSN:                  constant.PostgreConnection(),
	}
	if os.Getenv("APP_ENV") == "local" {
		configPostgre.DSN = constant.PostgreConnectionLocal()
	} else {
		configPostgre.DSN = constant.PostgreConnection()
	}
	return configPostgre
}

func GetJWTConfig() string {
	var jwt string
	if os.Getenv("APP_ENV") == "local" {
		jwt = os.Getenv("TEST_JWT_SECRET")
	} else {
		jwt = os.Getenv("JWT_SECRET")
	}
	return jwt
}
