package config

import (
	cli "gopkg.in/urfave/cli.v1"
	"k8s.io/client-go/kubernetes"
)

type Configuration struct {
	Client kubernetes.Interface
}

func (c *Configuration) AddFlags(app *cli.App) {
	flags := []cli.Flag{
		cli.StringFlag{},
	}
	app.Flags = append(app.Flags, flags...)
}
