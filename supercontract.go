package api

import (
	"context"

	idrive "github.com/proximax-storage/go-xpx-dfms-drive"
	sc "github.com/proximax-storage/go-xpx-dfms-drive/supercontract"
)

type SuperContract interface {
	Deploy(ctx context.Context, id idrive.ID, file string) error
	Execute(ctx context.Context, id sc.ID, gas uint64, function sc.Function) error
	
	Get(context.Context, sc.ID) (*sc.SuperContract, error)
	List(context.Context, idrive.ID) ([]sc.ID, error)
}
