// Package esign provides a simple interface genarate credentials
// for esign api service from a keyring entry
package esign

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/jfcote87/esign"
	"github.com/jfcote87/oauth2"
	"github.com/jfcote87/secrets"
)

// Credential returns credential for esign calls
func Credential(ctx context.Context, key, apiUserName string, tk *oauth2.Token, u *esign.UserInfo) (esign.Credential, error) {
	def, err := secrets.Get(secrets.ServiceDocuSign, key)
	if err != nil {
		return nil, err
	}
	var cfg *esign.JWTConfig
	if err := json.NewDecoder(bytes.NewBufferString(def)).Decode(&cfg); err != nil {
		return nil, err
	}
	return cfg.Credential(apiUserName, tk, u)
}
