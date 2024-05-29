package main

import (
	"reflect"
	"testing"
)

func Test_filterLines(t *testing.T) {
	type args struct {
		data  []string
		flags Flags
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Default",
			args: args{
				data: []string{"salam", "privet", "hello"},
				flags: Flags{
					Regexp: "privet",
				},
			},
			want:    []string{"privet"},
			wantErr: false,
		},
		{
			name: "Numeration",
			args: args{
				data: []string{"salam", "privet", "hello"},
				flags: Flags{
					Numeration: true,
					Regexp:     "privet",
				},
			},
			want:    []string{"1\tprivet"},
			wantErr: false,
		},
		{
			name: "Regexp",
			args: args{
				data: []string{"salam", "privet", "hello"},
				flags: Flags{
					Regexp: "((^|, )(privet|salam))+$",
				},
			},
			want:    []string{"salam", "privet"},
			wantErr: false,
		},
		{
			name: "Fixed",
			args: args{
				data: []string{"aboba 1", "aboba", "aboba 2"},
				flags: Flags{
					Fixed:  true,
					Regexp: "aboba",
				},
			},
			want:    []string{"aboba"},
			wantErr: false,
		},
		{
			name: "After",
			args: args{
				data: []string{"before", "test", "after"},
				flags: Flags{
					Fixed:  true,
					After:  1,
					Regexp: "test",
				},
			},
			want:    []string{"test", "after"},
			wantErr: false,
		},
		{
			name: "Before",
			args: args{
				data: []string{"before", "test", "after"},
				flags: Flags{
					Before: 1,
					Regexp: "test",
				},
			},
			want:    []string{"before", "test"},
			wantErr: false,
		},
		{
			name: "Context",
			args: args{
				data: []string{"aboba 1", "aboba", "aboba 2"},
				flags: Flags{
					Context: 1,
					Regexp:  "aboba",
				},
			},
			want:    []string{"aboba 1", "aboba", "aboba 2"},
			wantErr: false,
		},
		{
			name: "Invert",
			args: args{
				data: []string{"test 1", "test0", "test 2"},
				flags: Flags{
					Invert: true,
					Regexp: "test0",
				},
			},
			want:    []string{"test 1", "test 2"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := filterLines(tt.args.data, tt.args.flags)
			if (err != nil) != tt.wantErr {
				t.Errorf("filterLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterLines() got = %v, want %v", got, tt.want)
			}
		})
	}
}
