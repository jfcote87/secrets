// Package certify provides a simple interface for generating a certify
// api service from a keyring entry
package certify

import (
	"context"
	"encoding/json"

	"github.com/jfcote87/certify"
	"github.com/jfcote87/secrets"
)

// Service return service defined by key
func Service(ctx context.Context, key string) (*certify.Service, error) {
	def, err := secrets.Get(secrets.ServiceCertify, key)
	if err != nil {
		return nil, err
	}
	var cfg = make(map[string]string)
	if err = json.Unmarshal([]byte(def), &cfg); err != nil {
		return nil, err
	}
	return &certify.Service{
		Cred: certify.NewKeySecret(cfg["certifyKey"], cfg["certifySecret"]),
	}, nil
}
