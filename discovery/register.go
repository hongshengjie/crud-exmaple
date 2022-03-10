package discovery

import (
	"context"
	"fmt"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"go.etcd.io/etcd/client/v3/naming/resolver"

	"google.golang.org/grpc"
)

var dis *Discovery

type Discovery struct {
	Cli *clientv3.Client
	Em  endpoints.Manager
}

func init() {
	dis, _ = New("")
}

func Dis() *Discovery {
	return dis
}

func New(etcdaddress string) (*Discovery, error) {
	d := &Discovery{}
	var err error
	if etcdaddress == "" {
		etcdaddress = "http://localhost:2379"
	}
	d.Cli, err = clientv3.NewFromURL(etcdaddress)
	if err != nil {
		return nil, err
	}
	if d.Em, err = endpoints.NewManager(d.Cli, "crud-example.user"); err != nil {
		return nil, err
	}
	return d, nil
}

func Register(ctx context.Context, serviceID, instanceID, endpoint string) error {
	var err error
	if dis.Em == nil {
		if dis.Em, err = endpoints.NewManager(dis.Cli, serviceID); err != nil {
			return err
		}
	}
	lease := clientv3.NewLease(dis.Cli)
	tick, err := lease.Grant(ctx, 30)
	if err != nil {
		return err
	}
	ch, err := lease.KeepAlive(ctx, tick.ID)
	if err != nil {
		return err
	}
	go func() {
		for v := range ch {
			fmt.Println("lease ", v)
		}
	}()
	return dis.Em.AddEndpoint(ctx, instanceID, endpoints.Endpoint{Addr: endpoint}, clientv3.WithLease(tick.ID))
}

func DeleteRegister(ctx context.Context, instanceID string) error {
	if dis.Em != nil {
		return dis.Em.DeleteEndpoint(ctx, instanceID)
	}
	return nil

}

func NewConn(serviceID string) (*grpc.ClientConn, error) {
	resolver, err := resolver.NewBuilder(dis.Cli)
	if err != nil {
		return nil, err
	}
	return grpc.Dial(fmt.Sprintf("etcd:///%s", serviceID), grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`), grpc.WithInsecure(), grpc.WithResolvers(resolver))
}
