package ipvs

import (
	"flag"
	"fmt"

	log "github.com/zoumo/logdog"
)

func run(opt *Options) error {

	log.Infof("ipvs provider information %v\n", Version)
	return nil
}

func main() {
	flag.CommandLine.Parse([]string{})

	app := cli.NewApp()
	app.Name = "ingress sidecar"
	app.Compiled = time.Now()

}
