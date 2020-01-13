package api

import (
	"crypto/rand"
	"io"
	"io/ioutil"
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/multiformats/go-multihash"
)

func RandVerifyResult(t *testing.T) VerifyResult {
	return VerifyResult{
		{
			Replicator: randKey(t),
			FaultyBlocks: []cid.Cid{
				randCID(t),
			},
		},
		{
			Replicator: randKey(t),
			FaultyBlocks: []cid.Cid{
				randCID(t),
				randCID(t),
			},
		},
	}
}

func randKey(t *testing.T) crypto.PubKey {
	_, key, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}

	return key
}

func randCID(t *testing.T) cid.Cid {
	b, err := ioutil.ReadAll(io.LimitReader(rand.Reader, 256))
	if err != nil {
		t.Fatal(err)
	}

	hash, err := multihash.Sum(b, multihash.SHA2_256, -1)
	if err != nil {
		t.Fatal(err)
	}

	return cid.NewCidV1(cid.Raw, hash)
}
