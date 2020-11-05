// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package vrrp

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
	vpe "github.com/edwarnicke/govpp/pkg/v2005/binapi/vpe"
)

// RPCService defines RPC service  vrrp.
type RPCService interface {
	VrrpVrAddDel(ctx context.Context, in *VrrpVrAddDel) (*VrrpVrAddDelReply, error)
	VrrpVrDump(ctx context.Context, in *VrrpVrDump) (RPCService_VrrpVrDumpClient, error)
	VrrpVrPeerDump(ctx context.Context, in *VrrpVrPeerDump) (RPCService_VrrpVrPeerDumpClient, error)
	VrrpVrSetPeers(ctx context.Context, in *VrrpVrSetPeers) (*VrrpVrSetPeersReply, error)
	VrrpVrStartStop(ctx context.Context, in *VrrpVrStartStop) (*VrrpVrStartStopReply, error)
	VrrpVrTrackIfAddDel(ctx context.Context, in *VrrpVrTrackIfAddDel) (*VrrpVrTrackIfAddDelReply, error)
	VrrpVrTrackIfDump(ctx context.Context, in *VrrpVrTrackIfDump) (RPCService_VrrpVrTrackIfDumpClient, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) VrrpVrAddDel(ctx context.Context, in *VrrpVrAddDel) (*VrrpVrAddDelReply, error) {
	out := new(VrrpVrAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) VrrpVrDump(ctx context.Context, in *VrrpVrDump) (RPCService_VrrpVrDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_VrrpVrDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_VrrpVrDumpClient interface {
	Recv() (*VrrpVrDetails, error)
	api.Stream
}

type serviceClient_VrrpVrDumpClient struct {
	api.Stream
}

func (c *serviceClient_VrrpVrDumpClient) Recv() (*VrrpVrDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *VrrpVrDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) VrrpVrPeerDump(ctx context.Context, in *VrrpVrPeerDump) (RPCService_VrrpVrPeerDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_VrrpVrPeerDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_VrrpVrPeerDumpClient interface {
	Recv() (*VrrpVrPeerDetails, error)
	api.Stream
}

type serviceClient_VrrpVrPeerDumpClient struct {
	api.Stream
}

func (c *serviceClient_VrrpVrPeerDumpClient) Recv() (*VrrpVrPeerDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *VrrpVrPeerDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) VrrpVrSetPeers(ctx context.Context, in *VrrpVrSetPeers) (*VrrpVrSetPeersReply, error) {
	out := new(VrrpVrSetPeersReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) VrrpVrStartStop(ctx context.Context, in *VrrpVrStartStop) (*VrrpVrStartStopReply, error) {
	out := new(VrrpVrStartStopReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) VrrpVrTrackIfAddDel(ctx context.Context, in *VrrpVrTrackIfAddDel) (*VrrpVrTrackIfAddDelReply, error) {
	out := new(VrrpVrTrackIfAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) VrrpVrTrackIfDump(ctx context.Context, in *VrrpVrTrackIfDump) (RPCService_VrrpVrTrackIfDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_VrrpVrTrackIfDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_VrrpVrTrackIfDumpClient interface {
	Recv() (*VrrpVrTrackIfDetails, error)
	api.Stream
}

type serviceClient_VrrpVrTrackIfDumpClient struct {
	api.Stream
}

func (c *serviceClient_VrrpVrTrackIfDumpClient) Recv() (*VrrpVrTrackIfDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *VrrpVrTrackIfDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}
