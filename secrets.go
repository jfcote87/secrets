// Package secrets provides a simple interface to return
// credentials/key/dsn from keyring
package secrets

import (
	"io/ioutil"

	"github.com/zalando/go-keyring"
)

const (
	// ServiceSQL is id for sql dsn secrets
	ServiceSQL ServiceName = "sql"
	// ServiceSSHDB is id fro sshdb configuration secrets
	ServiceSSHDB ServiceName = "sshdb"
	// ServiceIntacct is id for intacct configuration secrets
	ServiceIntacct ServiceName = "intacct"
	// ServiceSalesforce is id for salesforce configurations
	ServiceSalesforce ServiceName = "sf"
	// ServiceCertify is id for certify configs
	ServiceCertify ServiceName = "certify"
	// ServiceDocuSign is id for docusign configs
	ServiceDocuSign ServiceName = "esign"
	// ServiceJira is id for jira jwt configs
	ServiceJira ServiceName = "jira"
	// ServiceGoogle is id for google jwt configs
	ServiceGoogle ServiceName = "google"
)

// ServiceName provides name for keyring list
type ServiceName string

func (a ServiceName) kr() string {
	return "kr_" + string(a)
}

// Get returns a secret from keyring
func Get(service ServiceName, key string) (string, error) {
	return keyring.Get(service.kr(), key)
}

// Set adds a new secret to the keyring
func Set(service ServiceName, key, value string) error {
	keyring.Set(service.kr(), key, value)
	return nil
}

// SetFromFile adds a new secret to the keyring
func SetFromFile(svnm ServiceName, key, filename string) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return Set(svnm, key, string(b))
}
