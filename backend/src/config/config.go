package config

import (
	"os"

	"github.com/joho/godotenv"
)

var ConnectionString string
var PORT string
var SECRET_KEY string

func LoadEnvVars() error {

	if err := godotenv.Load(); err != nil {
		return err
	}

	ConnectionString = os.Getenv("CONNECTION_STR")
	PORT = os.Getenv("PORT")
	SECRET_KEY = os.Getenv("SECRET_KEY")

	return nil

}
