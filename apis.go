package api

import "context"

type Node interface {
	// Network returns implementation of Network api
	Network() Network

	// Type returns type of node
	Type() NodeType

	// Versions returns a version of API
	Version(ctx context.Context) (*Versions, error)
}

// api.Client is scope of different apis available for DFMS Client node
type Client interface {
	Node

	// Contract return implementation of ContractClient api
	Contract() ContractClient

	// FS return implementation of DriveFS api
	FS() DriveFS

	// SuperContract return implementation of SuperContract api
	SuperContract() SuperContract
}

// api.Replicator is scope of different apis available for DFMS Replicator node
type Replicator interface {
	Node

	// Contract return implementation of ContractReplicator api
	Contract() ContractReplicator
}
