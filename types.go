package api

import (
	"encoding/json"
	"fmt"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/crypto"
)

// Note: should be updated manually
const apiVersion = "0.5.1"

// DFMS node type
type NodeType string

const (
	ReplicatorType NodeType = "dfms-replicator"
	ClientType     NodeType = "dfms-client"
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

// NewVersion creates new Version and fill all fields
func NewVersion(build, commit, date, system, gv string) Version {
	return Version{
		API:       apiVersion,
		Build:     build,
		Commit:    commit,
		BuildDate: date,
		System:    system,
		GO:        gv,
	}
}

type Version struct {
	// Current API version
	API string

	// Current Build version
	Build string

	// Commit hash
	Commit string

	// Data of node build
	BuildDate string

	// System for which the build was made
	System string

	// Version of Golang
	GO string
}

func (v Version) Marshal() ([]byte, error) {
	return json.Marshal(v)
}

func (v *Version) Unmarshal(data []byte) error {
	return json.Unmarshal(data, v)
}

func (v *Version) String() string {
	return fmt.Sprintf(
		"API Version: %s\n"+
			"Build Version: %s\n"+
			"Build date: %s\n"+
			"Commit: %s\n"+
			"System: %s\n"+
			"Golang version: %s\n",
		v.API,
		v.Build,
		v.Commit,
		v.BuildDate,
		v.System,
		v.GO,
	)
}
