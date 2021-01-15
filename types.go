package api

import (
	"encoding/json"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/crypto"
)

// Note: should be updated manually
const apiVersion = "0.4.1"

// DFMS node type
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

// NewVersion creates new version with current ApiVersion and wanted BuildVersion
func NewVersion(buildVersion string) *Version {
	return &Version{
		ApiVersion:   apiVersion,
		BuildVersion: buildVersion,
	}
}

type Version struct {
	// Current version of API
	ApiVersion string

	// Current version of app build
	BuildVersion string
}

func (v Version) MarshalJSON() ([]byte, error) {
	return json.Marshal(v)
}

func (v *Version) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, v)
}
