package api

import (
	"context"
	idrive "github.com/proximax-storage/go-xpx-dfms-drive"
	sc "github.com/proximax-storage/go-xpx-dfms-drive/supercontract"
)

type SuperContract interface {
	Deploy(ctx context.Context, id idrive.ID, file string, functions []*sc.Function) error
}
