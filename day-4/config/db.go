package config

import "fmt"

const (
	DBUser     = "root"
	DBPassword = "P@ssw0rd"
	DBName     = "users"
	DBHost     = "localhost"
	DBPort     = "3306"
)

func GetMySQLConnectionString() string {
	database := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUser, DBPassword, DBHost, DBPort, DBName)

	return database
}
