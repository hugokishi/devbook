package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// APIURL - Request Api Url
	APIURL = ""
	// PORT - Port for web app running
	PORT = 0
	// HashKey - Used for authenticate cookie
	HashKey []byte
	// BlockKey - Encrypt the cookie
	BlockKey []byte
)

// Load - Loads as environment variables
func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	PORT, err = strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
