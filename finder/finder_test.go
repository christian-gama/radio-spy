package finder

import "testing"

func TestFindListeners(t *testing.T) {
	type args struct {
		pattern string
		html    string
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "Find listeners",
			args: args{
				pattern: `listeners (\d+) bla bla`,
				html:    `bla bla listeners 100 bla bla`,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindListeners(tt.args.pattern, tt.args.html); got != tt.want {
				t.Errorf("FindListeners() = %v, want %v", got, tt.want)
			}
		})
	}
}
