package api

import (
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/crypto"
)

type NodeType string

const (
	ReplicatorType NodeType = "replicator"
	ClientType     NodeType = "client"
)

type VerifyResult []struct {
	// Replicator which failed verification and was excluded from contract immediately
	Replicator crypto.PubKey

	// List of defected blocks on which replicator failed verification
	FaultyBlocks []cid.Cid
}

func (res VerifyResult) MarshalJSON() ([]byte, error) {
	return MarshalVerifyResultJSON(res)
}

func (res *VerifyResult) UnmarshalJSON(data []byte) error {
	out, err := UnmarshalVerifyResultJSON(data)
	if err != nil {
		return err
	}

	*res = out
	return nil
}
