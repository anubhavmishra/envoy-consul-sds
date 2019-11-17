package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

const (
	portFlag             = "port"
	version              = "0.0.1"
	consulDatacenterFlag = "consul-datacenter"
	consulHTTPAddrFlag   = "consul-http-addr"
	consulACLTokenFlag   = "consul-acl-token"
	consulSchemeFlag     = "consul-scheme"
	consulAllowStaleFlag = "consul-stale"

	defaultPort             = 8080
	defaultConsulDatacenter = "dc1"
	defaultConsulHTTPAddr   = "localhost:8500"
	defaultConsulScheme     = "http"
)

func main() {
	app := cli.NewApp()
	configureCli(app)
	app.Action = mainAction
	app.Run(os.Args)
}

func mainAction(c *cli.Context) error {
	conf, err := validateConfig(c)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	server, err := NewServer(conf)
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	return server.Start()
}

func configureCli(app *cli.App) {
	app.Name = "envoy-consul-sds"
	app.Usage = "Envoy Consul Service Discovery Service"
	app.Version = version
	app.Flags = []cli.Flag{
		&cli.IntFlag{
			Name:  fmt.Sprintf("%s, p", portFlag),
			Value: defaultPort,
			Usage: "The `port` to start the webserver on",
		},
		&cli.StringFlag{
			Name:  consulDatacenterFlag,
			Value: defaultConsulDatacenter,
			Usage: "The `datacenter` for consul",
		},
		&cli.StringFlag{
			Name:  consulHTTPAddrFlag,
			Value: defaultConsulHTTPAddr,
			Usage: "The `address` for consul http api",
		},
		&cli.StringFlag{
			Name:  consulACLTokenFlag,
			Usage: "The acl token for consul",
		},
		&cli.StringFlag{
			Name:  consulSchemeFlag,
			Value: defaultConsulScheme,
			Usage: "The scheme for consul",
		},
		&cli.BoolFlag{
			Name:  consulAllowStaleFlag,
			Usage: "Set stale parameter on consul service health queries",
		},
	}
	cli.AppHelpTemplate = `{{.Name}} - {{.Usage}}

usage: {{.HelpName}} [options]
{{if .VisibleFlags}}
options:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}{{if .Version}}
version: {{.Version}}{{end}}
`
}

func validateConfig(c *cli.Context) (*ServerConfig, error) {
	// set defaults for non-required flags if not specified
	var port = c.Int(portFlag)
	var consulDatacenter = c.String(consulDatacenterFlag)
	var consulHTTPAddr = c.String(consulHTTPAddrFlag)
	var consulACLToken = os.Getenv("CONSUL_HTTP_TOKEN")
	var consulScheme = c.String(consulSchemeFlag)
	var consulAllowStaleFlag = c.Bool(consulAllowStaleFlag)

	if !c.IsSet(consulHTTPAddrFlag) {
		consulHTTPAddr = os.Getenv("CONSUL_HTTP_ADDR")
	}
	if c.IsSet(consulACLTokenFlag) {
		consulACLToken = c.String(consulACLTokenFlag)
	}

	return &ServerConfig{
		port:             port,
		consulDatacenter: consulDatacenter,
		consulHTTPAddr:   consulHTTPAddr,
		consulACLToken:   consulACLToken,
		consulScheme:     consulScheme,
		consulAllowStale: consulAllowStaleFlag,
	}, nil
}
