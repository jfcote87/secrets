// Package salesforcejwt provides a simple way to generate a JWT assertion
// from a keyring entry
package salesforcejwt

import (
	"github.com/jfcote87/oauth2/cache"
	"github.com/jfcote87/salesforce"
	"github.com/jfcote87/salesforce/auth/jwt"
	"github.com/jfcote87/secrets"
)

// Service reads definitions from keyring and returns new Service
func Service(key string, tokenCache cache.TokenCache) (*salesforce.Service, error) {
	def, err := secrets.Get(secrets.ServiceSalesforce, key)
	if err != nil {
		return nil, err
	}
	return jwt.ServiceFromJSON([]byte(def), tokenCache)
}
