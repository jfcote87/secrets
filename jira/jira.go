// Package jira provides a client for accessing the Jira API.
package jira
/*
import (
	"encoding/json"

	"github.com/jfcote87/jira"
	"github.com/jfcote87/secrets"
)

// Credential returns a jira credential
func Credential(key string, userName string) (jira.Credential, error) {
	def, err := secrets.Get(secrets.ServiceJira, key)
	if err != nil {
		return nil, err
	}
	var cfg *jira.AppJWTConfig
	if err := json.Unmarshal([]byte(def), &cfg); err != nil {
		return nil, err
	}
	if userName > "" {
		cfg.Subject = userName
	}
	return cfg.Credential()
} */
