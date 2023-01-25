package radio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRadio_GetUrl(t *testing.T) {
	type fields struct {
		name      string
		url       string
		listeners uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get URL",
			fields: fields{
				name: "Rádio 1",
				url:  "http://www.radio1.com.br",
			},
			want: "http://www.radio1.com.br",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Radio{
				name:      tt.fields.name,
				url:       tt.fields.url,
				listeners: tt.fields.listeners,
			}

			assert.Equal(t, tt.want, r.GetUrl(), "expected %d, got %d", tt.want, r.GetUrl())
		})
	}
}

func TestRadio_GetName(t *testing.T) {
	type fields struct {
		name             string
		url              string
		listeners        uint32
		listenersPattern string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get name",
			fields: fields{
				name:             "Rádio 1",
				url:              "http://www.radio1.com.br",
				listenersPattern: "[0-9]+",
			},
			want: "Rádio 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Radio{
				name:      tt.fields.name,
				url:       tt.fields.url,
				listeners: tt.fields.listeners,
			}

			assert.Equal(t, tt.want, r.GetName(), "expected %d, got %d", tt.want, r.GetName())
		})
	}
}

func TestRadio_GetListeners(t *testing.T) {
	type fields struct {
		name             string
		url              string
		listeners        uint32
		listenersPattern string
	}

	tests := []struct {
		name   string
		fields fields
		want   uint32
	}{
		{
			name: "Get listeners",
			fields: fields{
				name:             "Rádio 1",
				url:              "http://www.radio1.com.br",
				listeners:        100,
				listenersPattern: "[0-9]+",
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Radio{
				name:      tt.fields.name,
				url:       tt.fields.url,
				listeners: tt.fields.listeners,
			}

			assert.Equal(t, tt.want, r.GetListeners(), "expected %d, got %d", tt.want, r.GetListeners())
		})
	}
}

func TestNewRadio(t *testing.T) {
	type args struct {
		name             string
		url              string
		listenersPattern string
	}
	tests := []struct {
		name   string
		args   args
		want   *Radio
		panics bool
	}{
		{
			name: "New radio",
			args: args{
				name:             "Rádio 1",
				url:              "http://www.radio1.com.br",
				listenersPattern: "([0-9]+)",
			},
			want: &Radio{
				name:             "Rádio 1",
				url:              "http://www.radio1.com.br",
				listeners:        0,
				listenersPattern: "([0-9]+)",
			},
			panics: false,
		},
		{
			name: "New radio with empty name",
			args: args{
				name:             "",
				url:              "http://www.radio1.com.br",
				listenersPattern: "([0-9]+)",
			},
			want: &Radio{
				name:             "",
				url:              "http://www.radio1.com.br",
				listeners:        0,
				listenersPattern: "([0-9]+)",
			},
			panics: true,
		},
		{
			name: "New radio with empty url",
			args: args{
				name:             "Rádio 1",
				url:              "",
				listenersPattern: "([0-9]+)",
			},
			want: &Radio{
				name:             "Rádio 1",
				url:              "",
				listeners:        0,
				listenersPattern: "([0-9]+)",
			},
			panics: true,
		},
		{
			name: "New radio with empty listeners pattern",
			args: args{
				name:             "Rádio 1",
				url:              "http://www.radio1.com.br",
				listenersPattern: "",
			},
			want: &Radio{
				name:             "Rádio 1",
				url:              "http://www.radio1.com.br",
				listeners:        0,
				listenersPattern: "/",
			},
			panics: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panics {
				assert.Panics(t, func() { NewRadio(tt.args.name, tt.args.url, tt.args.listenersPattern) })
			} else {
				assert.NotPanics(t, func() { NewRadio(tt.args.name, tt.args.url, tt.args.listenersPattern) })
				assert.Equal(
					t, tt.want,
					NewRadio(
						tt.args.name,
						tt.args.url,
						tt.args.listenersPattern,
					),

					"expected %d, got %d",
					tt.want,
					NewRadio(
						tt.args.name,
						tt.args.url,
						tt.args.listenersPattern),
				)
			}
		})
	}
}
