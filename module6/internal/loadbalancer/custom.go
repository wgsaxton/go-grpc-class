package loadbalancer

import (
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
)

const Name = "ab_testing"

// create a constructor to create the load balancer builder
func NewBuilder(groups map[string]string, defaultAddr string) balancer.Builder {
	return base.NewBalancerBuilder(Name, pickerBuilder{
		groups:      groups,
		defaultAddr: defaultAddr,
	}, base.Config{HealthCheck: true})
}

// implement picker builder to instantiate out picker object
type pickerBuilder struct {
	groups      map[string]string
	defaultAddr string
}

func (p pickerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {

}

// implement our picker
type picker struct {
}
