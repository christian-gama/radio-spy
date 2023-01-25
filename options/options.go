package options

type Options struct {
	ShowQuantity bool
}

func NewOptions(showQuantity bool) *Options {
	return &Options{
		ShowQuantity: showQuantity,
	}
}
