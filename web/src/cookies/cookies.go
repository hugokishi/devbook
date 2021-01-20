package cookies

import (
	"net/http"
	"web/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

// Load - Setup cookie based on environment variables
func Load() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

// Save - Register authentication informations
func Save(w http.ResponseWriter, ID, Token string) error {
	datas := map[string]string{
		"id":    ID,
		"token": Token,
	}

	datasEncrypted, err := s.Encode("datas", datas)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "datas",
		Value:    datasEncrypted,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}
