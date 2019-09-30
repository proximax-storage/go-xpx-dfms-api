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
	Limit int
}

type ComposeOpt func(*ComposeOpts)
