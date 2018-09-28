package s3

import (
	"strings"
	"sync"

	SDK "github.com/aws/aws-sdk-go/service/s3"

	"github.com/rav94/hybrid-public-cloud-go-wrapper/config/aws"
	"github.com/rav94/hybrid-public-cloud-go-wrapper/log"
)

const (
	serviceName = "s3"
)

// S3 has S3 client and bucket list.
type S3 struct {
	client   *SDK.S3
	endpoint string

	logger log.Logger
	prefix string

	bucketsMu sync.RWMutex
	buckets   map[string]*Bucket
}

// New returns initialized *S4.
func S3NewConn(conf aws.AWSEnvConfig) (*S3, error) {
	sess, err := conf.Session()
	if err != nil {
		return nil, err
	}

	cli := SDK.New(sess)
	svc := &S3{
		client:   cli,
		endpoint: cli.ClientInfo.Endpoint,
		logger:   log.DefaultLogger,
		prefix:   conf.DefaultPrefix,
		buckets:  make(map[string]*Bucket),
	}
	return svc, nil
}