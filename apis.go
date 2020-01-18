package api

// api.Client is scope of different apis available for DFMS Client node
type Client interface {
	// Contract return implementation of ContractClient api
	Contract() ContractClient

	// FS return implementation of DriveFS api
	FS() DriveFS

	// Network returns implementation of Network api
	Network() Network

	// SuperContract return implementation of SuperContract api
	SuperContract() SuperContract
}

// api.Replicator is scope of different apis available for DFMS Replicator node
type Replicator interface {
	// Contract return implementation of ContractReplicator api
	Contract() ContractReplicator

	// Network returns implementation of Network api
	Network() Network
}
