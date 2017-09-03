package main

import "strings"

type Service struct {
	// A list of hosts that make up the service.
	Hosts []Host `json:"hosts"`
}

type Host struct {
	// The IP address of the upstream host.
	IPAddress string `json:"ip_address"`

	// The port of the upstream host.
	Port int32 `json:"port"`

	Tags *Tags `json:"tags,omitempty"`
}

type Tags struct {
	// The optional zone of the upstream host. Envoy uses the zone
	// for various statistics and load balancing tasks.
	AZ string `json:"az,omitempty"`

	// The optional canary status of the upstream host. Envoy uses
	// the canary status for various statistics and load balancing
	// tasks.
	Canary bool `json:"canary,omitempty"`

	// The optional load balancing weight of the upstream host, in
	// the range 1 - 100. Envoy uses the load balancing weight in
	// some of the built in load balancers.
	LoadBalancingWeight int32 `json:"load_balancing_weight,omitempty"`
}

func getService(consul *Consul, serviceName string) (*Service, error) {
	hosts := make([]Host, 0)

	// extract service name from dns.
	// for example: we will get "redis" from "redis.service.dc1.consul"
	s := strings.Split(serviceName, ".")

	// get service address, ports and tags from consul
	consulService, err := consul.GetService(s[0])
	if err != nil {
		return &Service{Hosts: hosts}, err
	}
	if consulService == nil {
		return &Service{Hosts: hosts}, nil
	}
	// create service hosts
	for _, srv := range consulService {
		hosts = append(hosts, Host{IPAddress: srv.Address, Port: int32(srv.Port)})
	}

	return &Service{Hosts: hosts}, nil
}
