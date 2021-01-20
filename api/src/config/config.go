package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// DbConnectStr - String for connect in database
	DbConnectStr = ""

	// PORT - Port the application will run on
	PORT = 0

	// SecretKey - Key of jwt token
	SecretKey []byte
)

// Load - Loads all environment variables
func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		PORT = 3333
	}

	DbConnectStr = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)

	SecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
}
