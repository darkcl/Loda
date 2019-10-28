package matcher

import (
	"reflect"
	"testing"
)

func TestURLMatcher_Process(t *testing.T) {
	type fields struct {
		Matcher Matcher
	}
	type args struct {
		input string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want1  Matcher
	}{
		{
			name: "Matching URL",
			fields: fields{
				Matcher: &URLMatcher{},
			},
			args: args{
				input: "http://google.com",
			},
			want:  true,
			want1: nil,
		},
		{
			name: "Not a URL",
			fields: fields{
				Matcher: &URLMatcher{},
			},
			args: args{
				input: "foo",
			},
			want:  false,
			want1: nil,
		},
		{
			name: "Not a http / https url",
			fields: fields{
				Matcher: &URLMatcher{},
			},
			args: args{
				input: "postgres://user:pass@host.com:5432/path?k=v#f",
			},
			want:  false,
			want1: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := URLMatcher{
				Matcher: tt.fields.Matcher,
			}
			got, got1 := u.Process(tt.args.input)
			if got != tt.want {
				t.Errorf("URLMatcher.Process() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("URLMatcher.Process() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
