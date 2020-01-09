package api

import (
	"context"
	"errors"

	idrive "github.com/proximax-storage/go-xpx-dfms-drive"
)

var ErrAlreadyStarted = errors.New("accepting already started")

type ContractClient interface {
	// Compose synchronously announces invites to the Network with current node as an
	// owner and tries to find members which agrees on specified parameters and options. It does not
	// guarantee success on resolving members. On success persists contract locally and gives
	// ability to use Drive.
	Compose(ctx context.Context, space, duration uint64, opts ...ComposeOpt) (*idrive.Contract, error)

	// Lists all the contracts in which Node participates as an owner or member.
	List(context.Context) ([]idrive.ID, error)

	// Get searches for contract in local storage and/or in blockchain.
	Get(context.Context, idrive.ID) (*idrive.Contract, error)

	// Amendments create subscription for Drive contract corrections for contract in local storage and/or in blockchain.
	Amendments(context.Context, idrive.ID) (ContractSubscription, error)

	// Verify initiates verification round between replicators.
	Verify(context.Context, idrive.ID) (VerifyResult, error)

	// Finish contract
	Finish(context.Context, idrive.ID) error
}

type ContractReplicator interface {
	ContractClient

	// Accept joins contract by it's id.
	// Can join only contracts awaiting new members.
	Accept(context.Context, idrive.ID) error

	// Accepted returns subscription for accepted contracts by the node.
	Accepted(context.Context) (ContractSubscription, error)

	// Invites creates subscription for new contract invitations. Main use case is to have
	// external contract acceptance logic.
	Invites(context.Context) (InviteSubscription, error)
}

// InviteSubscription for DriveInvitations
type InviteSubscription interface {
	// Next waits and blocks till new Invitation is received
	Next(context.Context) (*idrive.Invite, error)

	// Cancel stops subscription, like context canceling
	Close() error
}

// ContractSubscription for Drive
type ContractSubscription interface {
	// Next waits and blocks till new Drive update received
	Next(context.Context) (*idrive.Contract, error)

	// Cancel stops subscription, like context canceling
	Close() error
}
