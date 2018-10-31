package controller

import (
	"fmt"
	glog "github.com/zoumo/logdog"

	lbpoolv1alpha1 "github.com/bretagne-peiqi/loadbalancer-pool/pkg/apis/lbpool/v1alpha1"
	clientset "github.com/bretagne-peiqi/loadbalancer-pool/pkg/client/clientset/versioned"
	lbpoolscheme "github.com/bretagne-peiqi/loadbalancer-pool/pkg/client/clientset/versioned/scheme"
	dinformers "github.com/bretagne-peiqi/loadbalancer-pool/pkg/client/informers/externalversions/lbpool/v1alpha1"
	listers "github.com/bretagne-peiqi/loadbalancer-pool/pkg/client/listers/lbpool/v1alpha1"
)

const controllerAgentName = "loadbalancer-controller"

const (
	SuccessSynced         = "Synced"
	ErrResourceExists     = "ErrResourceExists"
	MessageResourceExists = "Resource %q already exists and is not managed by Loadbalancer"
	MessageResourceSynced = "Loadbalancer synced successfully"
)

func controller() {
}
