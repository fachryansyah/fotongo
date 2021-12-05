package constant

import (
	"os"
	"strings"
)

func PostgreConnection() string {
	dbConfigs := []string{
		"host=" + os.Getenv("DB_HOST"),
		"dbname=" + os.Getenv("DB_NAME"),
		"user=" + os.Getenv("DB_USER"),
		"password=" + os.Getenv("DB_PASS"),
		"sslmode=disable"}

	return strings.Join(dbConfigs, " ")
}

func PostgreConnectionTest() string {
	dbConfigs := []string{
		"host=" + os.Getenv("TEST_DB_HOST"),
		"dbname=" + os.Getenv("TEST_DB_NAME"),
		"user=" + os.Getenv("TEST_DB_USER"),
		"password=" + os.Getenv("TEST_DB_PASS"),
		"sslmode=disable"}

	return strings.Join(dbConfigs, " ")
}

func PostgreConnectionLocal() string {
	dbConfigs := []string{
		"host=" + os.Getenv("DB_HOST"),
		"dbname=" + os.Getenv("DB_NAME"),
		"user=" + os.Getenv("DB_USER"),
		"password=" + os.Getenv("DB_PASS"),
		"sslmode=disable"}

	return strings.Join(dbConfigs, " ")
}
