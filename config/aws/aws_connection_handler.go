package aws

import (
	"github.com/pkg/errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/rav94/hybrid-public-cloud-go-wrapper/private/pointer"
)

const awsDefaultRegion = "us-west-2"

type AWSEnvConfig struct {
	AccessKey   string
	SecretKey   string
	Region      string
	Endpoint    string

	// Filename - Linux/OSX: "$HOME/.aws/credentials" Windows:   "%USERPROFILE%\.aws\credentials"
	// Profile are used for file credentials
	Filename    string
	Profile     string

	// DefaultPrefix is used for service resource prefix
	// e.g.) DynamoDB table, S3 bucket, SQS Queue
	DefaultPrefix string

	// Specific sevice's options
	S3ForcePathStyle bool
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
		Region:      pointer.String(c.getRegion()),
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
	return awsDefaultRegion
}