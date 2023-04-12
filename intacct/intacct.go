// Package intacct provides a client for accessing the Intacct API.
package intacct

import (
	"bytes"

	"github.com/jfcote87/intacct"
	"github.com/jfcote87/secrets"
)

// Service returns an intacct service defined by key contents
func Service(key string, opts ...intacct.ConfigOption) (*intacct.Service, error) {
	def, err := secrets.Get(secrets.ServiceIntacct, key)
	if err != nil {
		return nil, err
	}
	return intacct.ServiceFromConfigJSON(bytes.NewReader([]byte(def)), opts...)
}
