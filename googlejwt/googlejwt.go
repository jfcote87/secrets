// Package googlejwt provides a simple way to generate a Google JWT Tokensource
package googlejwt

import (
	"net/http"

	"github.com/jfcote87/oauth2/google"
	"github.com/jfcote87/secrets"
)

// Client returns a jwt config from secret
func Client(key string, userName string, scopes ...string) (*http.Client, error) {
	def, err := secrets.Get(secrets.ServiceGoogle, key)
	if err != nil {
		return nil, err
	}
	cfg, err := google.JWTConfigFromJSON([]byte(def), scopes...)
	if err != nil {
		return nil, err
	}
	if userName > "" {
		cfg.Subject = userName
	}
	return cfg.Client(nil)
}
