// sshdb creates an ssh database using credentials for the keyring
package sshdb

import (
	"bytes"
	"database/sql"
	"encoding/json"

	"github.com/jfcote87/secrets"
	"github.com/jfcote87/sshdb"
)

// Connections returns DBs from  from TunnelConfig
func Connections(driver sshdb.Driver, key string) ([]*sql.DB, error) {
	def, err := secrets.Get(secrets.ServiceSSHDB, key)
	if err != nil {
		return nil, err
	}
	var cfg *sshdb.Config
	if err := json.NewDecoder(bytes.NewBufferString(def)).Decode(&cfg); err != nil {
		return nil, err
	}
	return cfg.OpenDBs(driver)
}
