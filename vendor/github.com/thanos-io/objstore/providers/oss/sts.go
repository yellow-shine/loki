package oss

import (
	"fmt"
	"log"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/aliyun/credentials-go/credentials"
)

type stsCred struct {
	AccessKeyId     string
	AccessKeySecret string
	SecurityToken   string
}
func (s *stsCred) GetAccessKeyID() string     { return s.AccessKeyId }
func (s *stsCred) GetAccessKeySecret() string { return s.AccessKeySecret }
func (s *stsCred) GetSecurityToken() string   { return s.SecurityToken }

type provider struct{ cred credentials.Credential }

// Compile-time interface checks
var _ oss.CredentialsProvider = (*provider)(nil)
var _ oss.CredentialsProviderE = (*provider)(nil)

func (p *provider) GetCredentialsE() (oss.Credentials, error) {
	return p.getCredentials()
}

func (p *provider) GetCredentials() oss.Credentials {
	c, err := p.getCredentials()
	if err != nil {
		log.Fatalf("get sts credential: %v", err)
	}
	return c
}

func (p *provider) getCredentials() (oss.Credentials, error) {
	c, err := p.cred.GetCredential()
	if err != nil {
		return nil, fmt.Errorf("get sts credential: %v", err)
	}
	return &stsCred{
		AccessKeyId:     *c.AccessKeyId,
		AccessKeySecret: *c.AccessKeySecret,
		SecurityToken:   *c.SecurityToken,
	}, nil
}

