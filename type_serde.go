package api

import (
	"encoding/hex"
	"encoding/json"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/crypto"
)

func MarshalVerifyResultJSON(res VerifyResult) ([]byte, error) {
	out := make(verifyResultJson, len(res))
	for i, r := range res {
		out[i] = struct {
			Replicator   string
			FaultyBlocks []string
		}{}

		b, err := crypto.MarshalPublicKey(r.Replicator)
		if err != nil {
			return nil, err
		}
		out[i].Replicator = hex.EncodeToString(b)

		out[i].FaultyBlocks = make([]string, len(r.FaultyBlocks))
		for j, id := range r.FaultyBlocks {
			out[i].FaultyBlocks[j] = id.String()
		}
	}

	return json.Marshal(out)
}

func UnmarshalVerifyResultJSON(data []byte) (VerifyResult, error) {
	in := verifyResultJson{}
	err := json.Unmarshal(data, &in)
	if err != nil {
		return nil, err
	}

	out := make(VerifyResult, len(in))
	for i, r := range in {
		out[i] = struct {
			Replicator   crypto.PubKey
			FaultyBlocks []cid.Cid
		}{}

		b, err := hex.DecodeString(r.Replicator)
		if err != nil {
			return out, err
		}

		out[i].Replicator, err = crypto.UnmarshalPublicKey(b)
		if err != nil {
			return out, err
		}

		out[i].FaultyBlocks = make([]cid.Cid, len(r.FaultyBlocks))
		for j, id := range r.FaultyBlocks {
			out[i].FaultyBlocks[j], err = cid.Decode(id)
			if err != nil {
				return out, err
			}
		}
	}

	return out, nil
}

type verifyResultJson []struct {
	Replicator   string
	FaultyBlocks []string
}
