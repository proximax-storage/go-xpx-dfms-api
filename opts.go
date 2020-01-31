package api

import (
	"errors"
	"time"

	"github.com/libp2p/go-libp2p-core/crypto"
)

const (
	minSubPeriod = time.Minute

	replicasDefault             = uint16(3)
	minReplicasDefault          = uint16(3)
	percentApproversDefault     = uint8(66)
	numberBillingPeriodsDefault = int64(3)
)

var ErrInvalidDriveSpace = errors.New("drive space can't be 0 or less")
var ErrInvalidBillingPeriod = errors.New("billing period can't be less 1min")
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
	MinReplicators       uint16
	PercentApprovers     uint8
	NumberBillingPeriods int64
	BillingPrice         int64
	Replicas             uint16
	PrivateKey           crypto.PrivKey
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

func BillingPrice(billingPrice int64) ComposeOpt {
	return func(opts *composeOpts) {
		if billingPrice > 0 {
			opts.BillingPrice = billingPrice
		}
	}
}

func NumberBillingPeriods(numberBillingPeriods int64) ComposeOpt {
	return func(opts *composeOpts) {
		if numberBillingPeriods > 0 {
			opts.NumberBillingPeriods = numberBillingPeriods
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
func Parse(space uint64, billingPeriod time.Duration, options ...ComposeOpt) (*composeOpts, error) {
	opts := &composeOpts{
		Replicas:             replicasDefault,
		MinReplicators:       minReplicasDefault,
		NumberBillingPeriods: numberBillingPeriodsDefault,
		PercentApprovers:     percentApproversDefault,
	}
	for _, o := range options {
		o(opts)
	}

	return opts, validate(space, billingPeriod, opts)
}

//validate validates passed arguments and composeOpts
func validate(space uint64, subPeriod time.Duration, opts *composeOpts) error {
	if space <= 0 {
		return ErrInvalidDriveSpace
	}
	if subPeriod < minSubPeriod {
		return ErrInvalidBillingPeriod
	}
	if opts.BillingPrice <= 0 {
		opts.BillingPrice = int64(space) * int64(opts.Replicas)
	}
	if opts.MinReplicators > opts.Replicas {
		return ErrManyMinReplicators
	}
	return nil
}
