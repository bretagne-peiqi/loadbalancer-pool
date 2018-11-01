package nginx

import (
	"fmt"
	"github.com/bretagne-peiqi/loadbalancer-pool/pkg/config"
	cli "gopkg.in/urfave/cli.v1"
)

type Options struct {
	Kubeconfig string
	Debug      bool
	Config     config.Configuration
}

func NewOptions() *Options {
	return &Options{}
}

func (opts *Options) AddFlags(app *cli.App) {
	opts.Config.AddFlags(app)

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "kubeconfig",
			Usage: "Path to a kube config",
		},
		cli.BoolFlag{
			Name:        "debug",
			Usage:       "Run with debug mode",
			Destination: &opts.Debug,
		},
	}

	app.Flags = append(app.Flags, flags...)
}
