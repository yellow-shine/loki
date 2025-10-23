package cos

import (
	"flag"

	"github.com/grafana/dskit/flagext"
)

// Config holds the configuration for Tencent Cloud COS client
type Config struct {
	Endpoint        string         `yaml:"endpoint"`
	Bucket          string         `yaml:"bucket"`
	Region          string         `yaml:"region"`
	AccessKeyID     string         `yaml:"access_key_id"`
	AccessKeySecret flagext.Secret `yaml:"access_key_secret"`
	AppId           string         `yaml:"app_id"`
}

// RegisterFlags registers the flags for Tencent Cloud COS storage config
func (cfg *Config) RegisterFlags(f *flag.FlagSet) {
	cfg.RegisterFlagsWithPrefix("", f)
}

// RegisterFlagsWithPrefix registers the flags for Tencent Cloud COS storage config with prefix
func (cfg *Config) RegisterFlagsWithPrefix(prefix string, f *flag.FlagSet) {
	f.StringVar(&cfg.Bucket, prefix+"cos.bucketname", "", "Name of cos bucket.")
	f.StringVar(&cfg.Endpoint, prefix+"cos.endpoint", "", "Endpoint to connect to.")
	f.StringVar(&cfg.AccessKeyID, prefix+"cos.access-key-id", "", "tencentcloud Access Key ID")
	f.Var(&cfg.AccessKeySecret, prefix+"cos.access-key-secret", "tencentcloud Secret Access Key")
	f.StringVar(&cfg.AppId, prefix+"cos.app_id", "", "tencentcloud App ID")
	f.StringVar(&cfg.Region, prefix+"cos.region", "", "tencentcloud Region")
}
