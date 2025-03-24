package cos

import (
	"github.com/go-kit/log"
	"github.com/thanos-io/objstore"
	"github.com/thanos-io/objstore/providers/cos"
)

// NewBucketClient creates a new Alibaba Cloud OSS bucket client
func NewBucketClient(cfg Config, component string, logger log.Logger) (objstore.Bucket, error) {
	cosCfg := cos.Config{
		Bucket:          cfg.Bucket,
		Region:          cfg.Region,
		Endpoint:        cfg.Endpoint,
		AppId:           cfg.AppId,
		SecretId:        cfg.AccessKeyID,
		SecretKey:       cfg.AccessKeySecret.String(),
	}
	return cos.NewBucketWithConfig(logger, cosCfg, component, nil)
}
