// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package lacp

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
	vpe "github.com/edwarnicke/govpp/pkg/v2005/binapi/vpe"
)

// RPCService defines RPC service  lacp.
type RPCService interface {
	SwInterfaceLacpDump(ctx context.Context, in *SwInterfaceLacpDump) (RPCService_SwInterfaceLacpDumpClient, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) SwInterfaceLacpDump(ctx context.Context, in *SwInterfaceLacpDump) (RPCService_SwInterfaceLacpDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_SwInterfaceLacpDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_SwInterfaceLacpDumpClient interface {
	Recv() (*SwInterfaceLacpDetails, error)
	api.Stream
}

type serviceClient_SwInterfaceLacpDumpClient struct {
	api.Stream
}

func (c *serviceClient_SwInterfaceLacpDumpClient) Recv() (*SwInterfaceLacpDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *SwInterfaceLacpDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}
