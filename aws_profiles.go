package main

import (
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	ini "gopkg.in/ini.v1"
)

var envSharedCredentialsFile = "AWS_SHARED_CREDENTIALS_FILE"
var envAWSConfigFile = "AWS_CONFIG_FILE"

// LoadProfile loads aws credentials
func LoadProfile() (*ini.File, error) {
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}
	sharedCredentialsPath, err := envOrDefault(envSharedCredentialsFile, func() (string, error) {
		return filepath.Join(home, ".aws", "credentials"), nil
	})
	configFilePath, err := envOrDefault(envAWSConfigFile, func() (string, error) {
		return filepath.Join(home, ".aws", "config"), nil
	})
	if err != nil {
		return nil, err
	}
	return ini.Load(sharedCredentialsPath, configFilePath)
}
