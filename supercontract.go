package api

import (
	"context"

	"github.com/ipfs/go-cid"
	idrive "github.com/proximax-storage/go-xpx-dfms-drive"
	sc "github.com/proximax-storage/go-xpx-dfms-drive/supercontract"
)

type SuperContract interface {
	Deploy(ctx context.Context, id idrive.ID, file string) (sc.ID, error)
	Execute(ctx context.Context, id sc.ID, gas uint64, function sc.Function) (cid.Cid, error)

	Get(context.Context, sc.ID) (*sc.SuperContract, error)
	List(context.Context, idrive.ID) ([]sc.ID, error)
	GetResults(ctx context.Context, id cid.Cid) ([]string, error)
	GetSuperContractExecutionsHash(ctx context.Context, id sc.ID) ([]cid.Cid, error)
}
