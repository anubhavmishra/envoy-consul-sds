package main

import (
	"github.com/hashicorp/consul/api"
	"github.com/pkg/errors"
)

type Consul struct {
	client    *api.Client
	queryOpts *api.QueryOptions
}

type ConsulService struct {
	Name    string
	Address string
	Port    int
}

func (c *Consul) GetService(serviceName string) ([]ConsulService, error) {
	serviceAddressesPorts := []ConsulService{}
	// get consul service addresses and ports
	addresses, _, err := c.client.Health().Service(serviceName, "", true, c.queryOpts)
	if err != nil {
		return nil, errors.Wrap(err, "get consul service")
	}

	for _, addr := range addresses {
		// append service addresses and ports
		serviceAddressesPorts = append(serviceAddressesPorts, ConsulService{
			Name:    addr.Service.Service,
			Address: addr.Node.Address,
			Port:    addr.Service.Port})
	}

	return serviceAddressesPorts, nil
}
