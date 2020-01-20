package api

import (
	"context"
	idrive "github.com/proximax-storage/go-xpx-dfms-drive"
)

type SuperContract interface {
	Deploy(ctx context.Context, id idrive.ID, file string) error
}
