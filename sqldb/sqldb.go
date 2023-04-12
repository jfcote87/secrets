// Package sqldb provides a simple way to get an
// mssql db connection from a keyring
package sqldb

import (
	"database/sql"
	"strings"

	"github.com/jfcote87/secrets"
)

// Open returns a sql db
func Open(key string) (*sql.DB, error) {
	def, err := secrets.Get(secrets.ServiceSQL, key)
	if err != nil {
		return nil, err
	}
	ty := strings.Split(key, "-")
	return sql.Open(ty[0], def)
}
