package resolve

import (
	"log"

	"google.golang.org/grpc/resolver"
)

var serverAddresses = []string{"localhost:50051", "localhost:50052", "localhost:50053"}

type Builder struct {
}

func (b Builder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &simpleResolver{cc: cc}
	r.Start()

	return r, nil
}

func (b Builder) Scheme() string {
	return "garrett"
}

type simpleResolver struct {
	cc resolver.ClientConn
}

func (s simpleResolver) Start() {
	addr := make([]resolver.Address, len(serverAddresses))
	for i, address := range serverAddresses {
		addr[i] = resolver.Address{Addr: address}
	}

	if err := s.cc.UpdateState(resolver.State{Addresses: addr}); err != nil {
		log.Fatal(err)
	}
}

func (s simpleResolver) ResolveNow(resolver.ResolveNowOptions) {}

func (s simpleResolver) Close() {}