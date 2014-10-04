package conoha

import (
	"fmt"
	"log"
	"os"

	"github.com/tkak/conoha"
)

type Config struct {
	Tenant   string `mapstructure:"tenant"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Token    string `mapstructure:"token"`
	Endpoint string `mapstructure:"endpoint"`
}

func (c *Config) Client() (*conoha.Client, error) {

	if v := os.Getenv("CONOHA_TENANT"); v != "" {
		c.Tenant = v
	}

	if v := os.Getenv("CONOHA_USER"); v != "" {
		c.User = v
	}

	if v := os.Getenv("CONOHA_PASSWORD"); v != "" {
		c.Password = v
	}

	client, err := conoha.NewClient(c.Tenant, c.User, c.Password)

	if err != nil {
		fmt.Println("")
	}

	log.Printf("[INFO] ConoHaClient configured for user %s", c.User)
	log.Printf("[INFO] ConoHa Endpoint configured %s", c.Endpoint)

	return client, nil
}
