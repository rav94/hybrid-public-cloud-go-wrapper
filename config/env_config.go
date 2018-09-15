package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/rav94/hybrid-public-cloud-go-wrapper/private/pointers"
)

type EnvConfig struct {
	PublicCloudProvider string
	AccessKey           string
	SecretKey           string
	Region    	    string
	Endpoint            string

	// Filename - Linux/OSX: "$HOME/.aws/credentials" Windows:   "%USERPROFILE%\.aws\credentials"
	// Profile are used for file credentials
	Filename string
	Profile  string
}


// Creating session to public cloud provider
func (c EnvConfig) CloudProviderSession() {
	if c.PublicCloudProvider == "AWS" {
		
	}
}