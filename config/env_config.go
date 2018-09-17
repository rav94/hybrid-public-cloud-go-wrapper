package config

import (
	"github.com/pkg/errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/rav94/hybrid-public-cloud-go-wrapper/private/pointers"
)

type EnvConfig struct {
	PublicCloudProvider string
}

type AWSEnvConfig struct {
	AccessKey   string
	SecretKey   string
	Region      string
	Endpoint    string

	// Filename - Linux/OSX: "$HOME/.aws/credentials" Windows:   "%USERPROFILE%\.aws\credentials"
	// Profile are used for file credentials
	Filename    string
	Profile     string
}

// Creating session to public cloud provider
func (c EnvConfig) CloudProviderSession(int, error) {
	if c.PublicCloudProvider != "AWS" {
		return -1, errors.New("Provided Public Cloud provider dosen't exist, please check")
	}

	switch c.PublicCloudProvider {
		case "AWS":
			

	}
}

// Session creates AWS session from the Config values.
func (c AWSEnvConfig) Session() (*session.Session, error) {
	return session.NewSession(c.AWSConfig())
}

// AWSConfig creates *aws.Config object from the fields.
func (c AWSEnvConfig) AWSConfig() *aws.Config {
	cred := c.awsCredentials()
	awsConf := &aws.Config{
		Credentials: cred,
		Region:      pointers.String(c.getRegion()),
	}

	return awsConf
}

func (c AWSEnvConfig) awsCredentials() *credentials.Credentials {
	// from env
	cred := credentials.NewEnvCredentials()
	_, err := cred.Get()
	if err == nil {
		return cred
	}

	// from param
	cred = credentials.NewStaticCredentials(c.AccessKey, c.SecretKey, "")
	_, err = cred.Get()
	if err == nil {
		return cred
	}

	// from local file
	return credentials.NewSharedCredentials(c.Filename, c.Profile)
}

func (c AWSEnvConfig) getRegion() string {
	if c.Region != "" {
		return c.Region
	}
	reg := EnvRegion()
	if reg != "" {
		return reg
	}
	return defaultRegion
}