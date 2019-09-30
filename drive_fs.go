package api

import (
	"context"
	"os"

	"github.com/ipfs/go-cid"
	files "github.com/ipfs/go-ipfs-files"

	idrive "github.com/proximax-storage/go-xpx-dfms-drive"
)

type DriveFS interface {
	// Add adds file or directory to the path of the contract
	// NOTE: Makes changes only locally. Synchronise with replicators by using Flush option
	Add(ctx context.Context, id idrive.ID, path string, file files.Node, opts ...DriveOption) (cid.Cid, error)

	// Get gets file or directory from the given path of the contract
	// NOTE: Fetches data from remote replicators if not stored locally
	Get(ctx context.Context, id idrive.ID, path string, opts ...DriveOption) (files.Node, error)

	// Remove removes the file or directory from the path
	// NOTE: Removes only reference to the file. Fully remove by using Clear option
	// NOTE: Makes changes only locally. Synchronise with replicators by using Flush option
	Remove(ctx context.Context, id idrive.ID, path string, opts ...DriveOption) error

	// TODO Open needs deep rethinking
	// Open opens FD of the file for further file modifications, even if the file is not stored locally
	// NOTE: Fetches data from remote replicators if not stored locally
	// NOTE: Makes changes only locally. Synchronise with replicators by using Flush() method
	// Open(ctx context.Context, id idrive.ID, path string, opts ...DriveOption) (mfs.FileDescriptor, error)

	// Move moves file or directory from the givens source path to the given destination path
	// Use also to rename file or directory
	// NOTE: Makes changes only locally. Synchronise with replicators by using Flush option
	Move(ctx context.Context, id idrive.ID, src string, dst string, opts ...DriveOption) error

	// Copy copies file or directory from the givens source path to the given destination path
	// It does not makes the full copy of the file or directory, it just copies the reference
	// NOTE: Makes changes only locally. Synchronise with replicators by using Flush option
	Copy(ctx context.Context, id idrive.ID, src string, dst string, opts ...DriveOption) error

	// MakeDir creates new directory on the given path
	// NOTE: Makes changes only locally. Synchronise with replicators by using Flush option
	MakeDir(ctx context.Context, id idrive.ID, path string, opts ...DriveOption) error

	// Ls returns information about the files and directories under the given path
	// NOTE: Fetches data from remote replicators if not stored locally
	Ls(ctx context.Context, id idrive.ID, path string) ([]os.FileInfo, error)

	// Stat returns information about the file or directory under the given path
	// NOTE: Fetches data from remote replicators if not stored locally
	Stat(ctx context.Context, id idrive.ID, path string) (os.FileInfo, error)

	// Flush pushes state of the local Drive to all replicators
	Flush(ctx context.Context, id idrive.ID) error
}
