package loadbalancer

import (
	"log"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/metadata"
)

const Name = "ab_testing"

// create a constructor to create the load balancer builder
func NewBuilder(groups map[string]string, defaultAddr string) balancer.Builder {
	return base.NewBalancerBuilder(Name, &pickerBuilder{
		groups:      groups,
		defaultAddr: defaultAddr,
	}, base.Config{HealthCheck: true})
}

// implement picker builder to instantiate out picker object
type pickerBuilder struct {
	groups      map[string]string // key: user group, value: IP address of server/back end
	defaultAddr string
}

func (p *pickerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	// check if there are server connections ready
	if len(info.ReadySCs) == 0 {
		return base.NewErrPicker(balancer.ErrNoSubConnAvailable)
	}

	// map sub connections
	scs := make(map[string]balancer.SubConn)
	for sc, inf := range info.ReadySCs {
		scs[inf.Address.Addr] = sc
	}

	// instantiate out picker
	return &picker{
		subConns:    scs,
		groups:      p.groups,
		defaultAddr: p.defaultAddr,
	}
}

// implement our picker
type picker struct {
	subConns    map[string]balancer.SubConn
	groups      map[string]string
	defaultAddr string
}

func (p picker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	// get the metadata from the request
	md, ok := metadata.FromOutgoingContext(info.Ctx)
	if !ok {
		log.Println("unable to get metadata from context, using default address")
		return p.defaultConn()
	}

	// get the value of the user-group header
	name := md.Get("user-group")
	if len(name) < 1 {
		log.Println("user group not specified in metadata, using default address")
		return p.defaultConn()
	}

	// check the group and retreive the relevant address
	addr, ok := p.groups[name[0]]
	if !ok {
		log.Println("group not in list, using default address")
		return p.defaultConn()
	}

	// map address to sub connection
	subConn, ok := p.subConns[addr]
	if !ok {
		log.Println("address is not in list of addresses, using default address")
		return p.defaultConn()
	}

	// return our pick result with our sub connection
	return balancer.PickResult{SubConn: subConn}, nil
}

func (p picker) defaultConn() (balancer.PickResult, error) {
	conn, ok := p.subConns[p.defaultAddr]
	if !ok {
		return balancer.PickResult{}, balancer.ErrNoSubConnAvailable
	}

	return balancer.PickResult{SubConn: conn}, nil
}
