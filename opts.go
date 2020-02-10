package api

import (
	"errors"
	"time"

	"github.com/libp2p/go-libp2p-core/crypto"
)

const (
	minSubPeriod                     = time.Minute
	replicasDefault                  = uint16(3)
	minReplicasDefault               = uint16(3)
	percentApproversDefault          = uint8(66)
	numberSubscriptionPeriodsDefault = int64(3)
)

var ErrInvalidDriveSpace = errors.New("drive space can't be 0 or less")
var ErrInvalidSubscriptionPeriod = errors.New("subscription period can't be less 1min")
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
	MinReplicators            uint16
	PercentApprovers          uint8
	NumberSubscriptionPeriods int64
	SubscriptionPrice         int64
	Replicas                  uint16
	PrivateKey                crypto.PrivKey
}

type ComposeOpt func(*composeOpts)

func MinReplicators(minReplicators uint16) ComposeOpt {
	return func(opts *composeOpts) {
		if minReplicators > 0 {
			opts.MinReplicators = minReplicators
		}
	}
}

func PercentApprovers(percentApprovers uint8) ComposeOpt {
	return func(opts *composeOpts) {
		if percentApprovers > 0 {
			opts.PercentApprovers = percentApprovers
		}
	}
}

func SubscriptionPrice(SubscriptionPrice int64) ComposeOpt {
	return func(opts *composeOpts) {
		if SubscriptionPrice > 0 {
			opts.SubscriptionPrice = SubscriptionPrice
		}
	}
}

func NumberSubscriptionPeriods(numberSubscriptionPeriods int64) ComposeOpt {
	return func(opts *composeOpts) {
		if numberSubscriptionPeriods > 0 {
			opts.NumberSubscriptionPeriods = numberSubscriptionPeriods
		}
	}
}

func Replicas(replicas uint16) ComposeOpt {
	return func(opts *composeOpts) {
		if replicas > 0 {
			opts.Replicas = replicas
		}
	}
}

func PrivateKey(pk crypto.PrivKey) ComposeOpt {
	return func(opts *composeOpts) {
		opts.PrivateKey = pk
	}
}

// Parse parses the given options and return composeOpts
func Parse(space uint64, subPeriod time.Duration, options ...ComposeOpt) (*composeOpts, error) {
	opts := &composeOpts{
		Replicas:                  replicasDefault,
		MinReplicators:            minReplicasDefault,
		NumberSubscriptionPeriods: numberSubscriptionPeriodsDefault,
		PercentApprovers:          percentApproversDefault,
	}
	for _, o := range options {
		o(opts)
	}

	return opts, validate(space, subPeriod, opts)
}

//validate validates passed arguments and composeOpts
func validate(space uint64, subPeriod time.Duration, opts *composeOpts) error {
	if space <= 0 {
		return ErrInvalidDriveSpace
	}
	if subPeriod < minSubPeriod {
		return ErrInvalidSubscriptionPeriod
	}
	if opts.SubscriptionPrice <= 0 {
		opts.SubscriptionPrice = int64(space) * int64(opts.Replicas)
	}
	if opts.MinReplicators > opts.Replicas {
		return ErrManyMinReplicators
	}
	return nil
}
