package options

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOptions(t *testing.T) {
	type args struct {
		showQuantity bool
		server       bool
		port         int
	}
	tests := []struct {
		name   string
		args   args
		want   *Options
		panics bool
	}{
		{
			name: "New options",
			args: args{
				showQuantity: true,
				server:       true,
				port:         8080,
			},
			want: &Options{
				ShowQuantity: true,
				Server:       true,
				Port:         8080,
			},
			panics: false,
		},
		{
			name: "New options with invalid port",
			args: args{
				showQuantity: true,
				server:       true,
				port:         0,
			},
			want: &Options{
				ShowQuantity: true,
				Server:       true,
				Port:         0,
			},
			panics: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panics {
				assert.Panics(t, func() { NewOptions(tt.args.showQuantity, tt.args.server, tt.args.port) })
			} else {
				assert.NotPanics(t, func() { NewOptions(tt.args.showQuantity, tt.args.server, tt.args.port) })
				assert.Equal(
					t, tt.want,
					NewOptions(tt.args.showQuantity, tt.args.server, tt.args.port),

					"expected %d, got %d",
					tt.want, NewOptions(tt.args.showQuantity, tt.args.server, tt.args.port),
				)
			}
		})
	}
}
