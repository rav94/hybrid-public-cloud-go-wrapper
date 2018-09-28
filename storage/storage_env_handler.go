package storage

import (
	"github.com/pkg/errors"
	"github.com/rav94/hybrid-public-cloud-go-wrapper/config/aws"
	"github.com/rav94/hybrid-public-cloud-go-wrapper/storage/s3"
)

type StorageConfig struct {
	PublicCloudProvider string
	StorageSolutionType string
	EnvConfigs 	    map[string]interface{} //Gives ability to store any type as value
}

// Hold config data returned from different storage providers
type StorageData struct {

}


func NewStorageSession(c StorageConfig) (*StorageData, error) {
	if c.PublicCloudProvider != "AWS" || c.StorageSolutionType != "S3" || len(c.EnvConfigs == 0) {
		return nil, errors.New("Provided Configs seems to mismatch, please check")
	}

	switch c.PublicCloudProvider {
		case "AWS":
			svc, err := s3.S3NewConn(aws.AWSEnvConfig {
				AccessKey: c.EnvConfigs["AccessKey"],
				SecretKey: c.EnvConfigs["SecretKey"],
				Region: c.EnvConfigs["Region"],
				S3ForcePathStyle: c.EnvConfigs["S3ForcePathStyle"],
				Endpoint: c.EnvConfigs["Endpoint"],
			})

			if err != nil {
				return nil, errors.New("Error in creating new S3 client")
			}

			&StorageData{} = svc

			return &StorageData{}, nil
		case "GCP":
		case "AZURE":
	}


}