package nginx

import (
	"flag"
	"fmt"

	log "github.com/zoumo/logdog"
)

func run(opt *Options) error {

	log.Infof("nginx provider information %v\n")
	return nil
}

func main() {
	flag.CommandLine.Parse([]string{})

	app := cli.NewApp()
	app.Name = "ipvs sidecar"
	app.Compiled = time.Now()

}
