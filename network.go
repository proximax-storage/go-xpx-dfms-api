package api

import (
	"context"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
)

type Network interface {
	// Connect establishes connection to given multiaddr/s
	Connect(context.Context, ...multiaddr.Multiaddr) error

	// Disconnect destroys connection to given multiaddr/s
	Disconnect(context.Context, ...multiaddr.Multiaddr) error

	// Peers lists information about all connected peers
	Peers(context.Context) ([]*peer.AddrInfo, error)

	// ID returns node's identifier
	ID(context.Context) (peer.ID, error)

	// Addrs returns list of addresses node uses for p2p communications
	Addrs(context.Context) ([]multiaddr.Multiaddr, error)
}
