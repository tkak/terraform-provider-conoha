package conoha

import (
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
)

type Config struct {
	Username   string
	Password   string
	TenantName string
	Token      string
	Endpoint   string
}

func (c *Config) NewClient() (*gophercloud.ProviderClient, error) {
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: "https://ident-r1nd1001.cnode.jp/v2.0/",
		Username:         c.Username,
		Password:         c.Password,
		TenantName:       c.TenantName,
	}
	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		return nil, err
	}

	return provider, nil
}
