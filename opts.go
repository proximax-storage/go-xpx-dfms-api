package api

import (
	"errors"

	"github.com/libp2p/go-libp2p-core/crypto"
)

const (
	replicasDefault         = uint16(3)
	minReplicasDefault      = uint16(3)
	percentApproversDefault = uint8(66)
	billingPeriodDefault    = int64(172800)
)

var ErrInvalidDriveSpace = errors.New("drive Space can't be 0 or less")
var ErrInvalidDuration = errors.New("duration can't be 0 or less")
var ErrInvalidReplicas = errors.New("count of replicas can't be 0 or less")
var ErrInvalidBillingPeriod = errors.New("billing period can't be 0 or less")
var ErrInvalidBillingPeriodMultiply = errors.New("billing period is not a multiple of the duration")
var ErrInvalidPercentApprovers = errors.New("percent of approvers can't be 0 or less")
var ErrInvalidMinReplicators = errors.New("a minimum count of replicators can't be 0 or less")
var ErrManyMinReplicators = errors.New("a minimum count of replicators can't be more than count of replicas")

type DriveOption func(opts *DriveOptions)

type DriveOptions struct {
	Flush bool
	Clear bool
	Local bool
}

func Flush(f bool) DriveOption {
	return func(opts *DriveOptions) {
		opts.Flush = f
	}
}

func Clear(c bool) DriveOption {
	return func(opts *DriveOptions) {
		opts.Clear = c
	}
}

func Local(l bool) DriveOption {
	return func(opts *DriveOptions) {
		opts.Local = l
	}
}

func ParseDriveOptions(opts ...DriveOption) *DriveOptions {
	do := &DriveOptions{
		Flush: false,
		Clear: false,
	}

	for _, opt := range opts {
		opt(do)
	}

	return do
}

type composeOpts struct {
	MinReplicators   uint16
	PercentApprovers uint8
	BillingPeriod    int64
	BillingPrice     int64
	Replicas         uint16
	PrivateKey       crypto.PrivKey
}

type ComposeOpt func(*composeOpts)

func MinReplicators(minReplicators uint16) ComposeOpt {
	return func(opts *composeOpts) {
		opts.MinReplicators = minReplicators
	}
}

func PercentApprovers(percentApprovers uint8) ComposeOpt {
	return func(opts *composeOpts) {
		opts.PercentApprovers = percentApprovers
	}
}

func BillingPrice(billingPrice int64) ComposeOpt {
	return func(opts *composeOpts) {
		opts.BillingPrice = billingPrice
	}
}

func BillingPeriod(billingPeriod int64) ComposeOpt {
	return func(opts *composeOpts) {
		opts.BillingPeriod = billingPeriod
	}
}

func Replicas(replicas uint16) ComposeOpt {
	return func(opts *composeOpts) {
		opts.Replicas = replicas
	}
}

func PrivateKey(pk crypto.PrivKey) ComposeOpt {
	return func(opts *composeOpts) {
		opts.PrivateKey = pk
	}
}

// Parse parses the given options and return composeOpts
func Parse(space, duration uint64, options ...ComposeOpt) (*composeOpts, error) {
	opts := &composeOpts{
		Replicas:         replicasDefault,
		MinReplicators:   minReplicasDefault,
		BillingPeriod:    billingPeriodDefault,
		PercentApprovers: percentApproversDefault,
	}
	for _, o := range options {
		o(opts)
	}

	return opts, validate(space, duration, opts)
}

func validate(space, duration uint64, opts *composeOpts) error {
	if opts.BillingPeriod <= 0 {
		return ErrInvalidBillingPeriod
	}
	if opts.Replicas <= 0 {
		return ErrInvalidReplicas
	}
	if opts.BillingPrice <= 0 {
		opts.BillingPrice = int64(space) * int64(opts.Replicas)
	}
	if opts.MinReplicators <= 0 {
		return ErrInvalidMinReplicators
	}
	if opts.PercentApprovers <= 0 {
		return ErrInvalidPercentApprovers
	}
	if opts.MinReplicators > opts.Replicas {
		return ErrManyMinReplicators
	}
	if space <= 0 {
		return ErrInvalidDriveSpace
	}
	if duration <= 0 {
		return ErrInvalidDuration
	}
	if duration%uint64(opts.BillingPeriod) != 0 {
		return ErrInvalidBillingPeriodMultiply
	}
	return nil
}
