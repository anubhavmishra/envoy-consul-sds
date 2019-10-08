package main

import (
	"fmt"

	"log"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

// Server that serves kuberdbs web service
type Server struct {
	port    int
	version string
	engine  *gin.Engine
	consul  *Consul
}

type ServerConfig struct {
	port             int
	consulDatacenter string
	consulHTTPAddr   string
	consulACLToken   string
	consulScheme     string
	consulAllowStale bool
}

func NewServer(config *ServerConfig) (*Server, error) {
	// create a new engine
	engine := gin.New()

	// create consul config
	consulConfig := &api.Config{
		Datacenter: config.consulDatacenter,
		Address:    config.consulHTTPAddr,
		Token:      config.consulACLToken,
		Scheme:     config.consulScheme,
	}

	// create consul client
	// Set CONSUL_HTTP_ADDR in the environment that would automatically be used by the client
	client, err := api.NewClient(consulConfig)
	if err != nil {
		return nil, errors.Wrap(err, "initializing consul client")
	}

	var queryOpts *api.QueryOptions
	if config.consulAllowStale {
		queryOpts = &api.QueryOptions{AllowStale: true}

	}

	return &Server{
		port:    config.port,
		version: version,
		engine:  engine,
		consul:  &Consul{client: client, queryOpts: queryOpts},
	}, nil
}

func (s *Server) Start() error {
	s.engine.Use(gin.Recovery(), gin.Logger())
	s.engine.GET("/", s.index)
	s.engine.GET("/v1/registration/:service_name", s.registration)
	log.Printf("envoy-consul-sds started - listening on port %v", s.port)
	if err := s.engine.Run(fmt.Sprintf(":%v", s.port)); err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	return nil
}

func (s *Server) registration(c *gin.Context) {
	// get service name from url param
	serviceName := c.Param("service_name")

	service, err := getService(s.consul, serviceName)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, service)
		return
	}

	c.JSON(http.StatusOK, service)
}

func (s *Server) index(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"name": "envoy-consul-sds", "version": s.version})
}
