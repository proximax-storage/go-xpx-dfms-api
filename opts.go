package api

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

type ComposeOpts struct {
	MinReplicators   uint16
	PercentApprovers uint8
	BillingPeriod    int64
	BillingPrice     int64
}

type ComposeOpt func(*ComposeOpts)

func MinReplicators(minReplicators uint16) ComposeOpt {
	return func(opts *ComposeOpts) {
		opts.MinReplicators = minReplicators
	}
}

func PercentApprovers(percentApprovers uint8) ComposeOpt {
	return func(opts *ComposeOpts) {
		opts.PercentApprovers = percentApprovers
	}
}

func BillingPrice(billingPrice int64) ComposeOpt {
	return func(opts *ComposeOpts) {
		opts.BillingPrice = billingPrice
	}
}
func BillingPeriod(billingPeriod int64) ComposeOpt {
	return func(opts *ComposeOpts) {
		opts.BillingPeriod = billingPeriod
	}
}

// Apply applies the given options to this DiscoveryOpts
func (opts *ComposeOpts) Apply(options ...ComposeOpt) error {
	for _, o := range options {
		o(opts)
	}
	return nil
}
