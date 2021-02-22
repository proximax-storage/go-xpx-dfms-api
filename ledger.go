package api

import (
	"context"

	drive "github.com/proximax-storage/go-xpx-dfms-drive"
)

type Ledger interface {
	// ListDrives returns all active Drives registered on the Ledger.
	// Note: Pagination will be introduced later on.
	ListDrives(context.Context) ([]drive.ID, error)
}
