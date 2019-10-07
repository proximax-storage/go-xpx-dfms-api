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
	Compose(ctx context.Context, space, duration uint64, opts ...ComposeOpt) (idrive.Contract, error)

	// Lists all the contracts in which Node participates as an owner or member.
	List(context.Context) ([]idrive.ID, error)

	// Get searches for contract in local storage and/or in blockchain.
	Get(context.Context, idrive.ID) (idrive.Contract, error)

	// Amendments create subscription for Drive contract corrections for contract in local storage and/or in blockchain.
	Amendments(context.Context, idrive.ID) (idrive.ContractSubscription, error)
}

type ContractReplicator interface {
	ContractClient

	// Accept joins contract by it's id.
	// Can join only contracts awaiting new members.
	Accept(context.Context, idrive.ID) error

	// Accepted returns subscription for accepted contracts by the node.
	Accepted(context.Context) (idrive.ContractSubscription, error)

	// Invites creates subscription for new contract invitations. Main use case is to have
	// external contract acceptance logic.
	Invites(context.Context) (idrive.InviteSubscription, error)

	// StartAccepting tells the node to receive ContractReplicator invitations and accept them with defined
	// strategy. Method will make some basic validation regarding the contract before accepting.
	StartAccepting(context.Context, AcceptStrategy) error

	// StopAccepting tells the node to stop receiving new DriveContracts.
	StopAccepting(context.Context) error
}

// AcceptStrategy is a predefined group of validators
type AcceptStrategy uint8

const (
	AcceptAll AcceptStrategy = 0
)
