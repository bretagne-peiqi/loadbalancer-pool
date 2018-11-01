package ipvsdr

import (
	"fmt"
	"net"
	"os"
	"syscall"
	"text/template"
	"time"

	log "github.com/zoumo/logdog"

	"k8s.io/kubernetes/pkg/util/iptables"
)

const (
	iptablesChain  = "LOADBALANCER-IPVS"
	keepalivedCfg  = "/etc/keepalived/keepalived.conf"
	keepalivedTmpl = "/root/keepalived.tmpl"

	acceptMark = 1
	dropMark   = 0
	mask       = "0x00000001"
)

type keepalived struct {
	useUnicast bool
	vips       []string
	tmpl       *template.Template
	ipt        iptables.Interface
}

func (k *keepalived) UpdateConfig(vss []virtualServer, neighors []ipmac, priority, vrid int) error {
	f, err := os.Create(keepalivedCfg)
	if err != nil {
		log.Errorf("Failed to create keepalived configfile: %v\n", err)
		return err
	}
	defer f.Close()
	return nil
}

func (k *keepalived) run() {

}

func (k *keepalived) Reload() error {
	log.Info("hot reload keepalived config")
}
