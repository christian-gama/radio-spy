package options

// Options is the struct that contains the options of the application
type Options struct {
	ShowQuantity bool
}

// NewOptions returns a new Options struct
func NewOptions(showQuantity bool) *Options {
	return &Options{
		ShowQuantity: showQuantity,
	}
}
