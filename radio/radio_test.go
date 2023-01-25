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

func TestRadio_SetUrl(t *testing.T) {
	type fields struct {
		name      string
		url       string
		listeners uint32
	}

	type args struct {
		url string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		panics bool
	}{
		{
			name: "Valid http URL",
			fields: fields{
				name: "Rádio 1",
				url:  "http://www.radio1.com.br",
			},
			args: args{
				url: "http://www.radio1.com.br",
			},
			panics: false,
		},
		{
			name: "Valid https URL",
			fields: fields{
				name: "Rádio 1",
				url:  "https://www.radio1.com.br",
			},
			args: args{
				url: "https://www.radio1.com.br",
			},
			panics: false,
		},
		{
			name: "Invalid URL without protocol",
			fields: fields{
				name: "Rádio 1",
				url:  "http://www.radio1.com.br",
			},
			args: args{
				url: "www.radio1.com.br",
			},
			panics: true,
		},
		{
			name: "Invalid URL with invalid protocol",
			fields: fields{
				name: "Rádio 1",
				url:  "http://www.radio1.com.br",
			},
			args: args{
				url: "ftp://www.radio1.com.br",
			},
			panics: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Radio{
				name:      tt.fields.name,
				url:       tt.fields.url,
				listeners: tt.fields.listeners,
			}

			if tt.panics {
				assert.Panics(t, func() { r.SetUrl(tt.args.url) })
			} else {
				assert.NotPanics(t, func() { r.SetUrl(tt.args.url) })
				assert.Equal(t, tt.args.url, r.GetUrl(), "expected %d, got %d", tt.args.url, r.GetUrl())
			}
		})
	}
}

func TestRadio_GetName(t *testing.T) {
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
			name: "Get name",
			fields: fields{
				name: "Rádio 1",
				url:  "http://www.radio1.com.br",
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

func TestRadio_SetName(t *testing.T) {
	type fields struct {
		name      string
		url       string
		listeners uint32
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		panics bool
	}{
		{
			name: "Valid name",
			fields: fields{
				name: "Rádio 1",
				url:  "http://www.radio1.com.br",
			},
			args: args{
				name: "Rádio 1",
			},
			panics: false,
		},
		{
			name: "Invalid name",
			fields: fields{
				name: "Rádio 1",
				url:  "http://www.radio1.com.br",
			},
			args: args{
				name: "",
			},
			panics: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Radio{
				name:      tt.fields.name,
				url:       tt.fields.url,
				listeners: tt.fields.listeners,
			}

			if tt.panics {
				assert.Panics(t, func() { r.SetName(tt.args.name) })
			} else {
				assert.NotPanics(t, func() { r.SetName(tt.args.name) })
				assert.Equal(t, tt.args.name, r.GetName(), "expected %d, got %d", tt.args.name, r.GetName())
			}
		})
	}
}

func TestRadio_GetListeners(t *testing.T) {
	type fields struct {
		name      string
		url       string
		listeners uint32
	}

	tests := []struct {
		name   string
		fields fields
		want   uint32
	}{
		{
			name: "Get listeners",
			fields: fields{
				name:      "Rádio 1",
				url:       "http://www.radio1.com.br",
				listeners: 100,
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

func TestRadio_SetListeners(t *testing.T) {
	type fields struct {
		name      string
		url       string
		listeners uint32
	}
	type args struct {
		listeners uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Set listeners",
			fields: fields{
				name:      "Rádio 1",
				url:       "http://www.radio1.com.br",
				listeners: 100,
			},
			args: args{
				listeners: 200,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Radio{
				name:      tt.fields.name,
				url:       tt.fields.url,
				listeners: tt.fields.listeners,
			}

			assert.NotPanics(t, func() { r.SetListeners(tt.args.listeners) })
			assert.Equal(t, tt.args.listeners, r.GetListeners(), "expected %d, got %d", tt.args.listeners, r.GetListeners())
		})
	}
}

func TestNewRadio(t *testing.T) {
	type args struct {
		name string
		url  string
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
				name: "Rádio 1",
				url:  "http://www.radio1.com.br",
			},
			want: &Radio{
				name:      "Rádio 1",
				url:       "http://www.radio1.com.br",
				listeners: 0,
			},
			panics: false,
		},
		{
			name: "New radio with empty name",
			args: args{
				name: "",
				url:  "http://www.radio1.com.br",
			},
			want: &Radio{
				name:      "",
				url:       "http://www.radio1.com.br",
				listeners: 0,
			},
			panics: true,
		},
		{
			name: "New radio with empty url",
			args: args{
				name: "Rádio 1",
				url:  "",
			},
			want: &Radio{
				name:      "Rádio 1",
				url:       "",
				listeners: 0,
			},
			panics: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panics {
				assert.Panics(t, func() { NewRadio(tt.args.name, tt.args.url) })
			} else {
				assert.NotPanics(t, func() { NewRadio(tt.args.name, tt.args.url) })
				assert.Equal(t, tt.want, NewRadio(tt.args.name, tt.args.url), "expected %d, got %d", tt.want, NewRadio(tt.args.name, tt.args.url))
			}
		})
	}
}
