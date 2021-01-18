package api

import (
	"encoding/json"
	"fmt"

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

// NewVersion creates new version with current API and wanted Build
func NewVersion(buildVersion string) *Versions {
	return &Versions{
		API:   apiVersion,
		Build: buildVersion,
	}
}

type Versions struct {
	// Current version of API
	API string

	// Current version of app build
	Build string
}

func (v Versions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v)
}

func (v *Versions) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, v)
}

func (v *Versions) String() string {
	return fmt.Sprintf(
		"API Version: %s\n"+
			"Build Version: %s\n",
		v.API,
		v.Build,
	)
}
