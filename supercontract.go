package api

import (
	"context"
	"github.com/ipfs/go-cid"
	idrive "github.com/proximax-storage/go-xpx-dfms-drive"
)

type Supercontract interface {
	Deploy(ctx context.Context, id idrive.ID, file cid.Cid, functions []string) error
}
