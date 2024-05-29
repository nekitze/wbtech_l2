package main

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	type args struct {
		lines []string
		flags Flags
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Default",
			args{
				lines: []string{"A", "C", "B", "D"},
				flags: Flags{},
			},
			[]string{"A", "B", "C", "D"},
		},

		{name: "SortedColumn",
			args: args{
				lines: []string{"A d", "B c", "C b", "D a"},
				flags: Flags{SortColumn: 1},
			},
			want: []string{"D a", "C b", "B c", "A d"},
		},

		{"Reversed",
			args{
				lines: []string{"A", "C", "B", "D"},
				flags: Flags{Reversed: true},
			},
			[]string{"D", "C", "B", "A"},
		},

		{"NumberSorted",
			args{
				lines: []string{"2", "5", "4", "1"},
				flags: Flags{NumberSorted: true},
			},
			[]string{"1", "2", "4", "5"},
		},

		{"Empty",
			args{
				lines: []string{},
				flags: Flags{},
			},
			[]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sort(tt.args.lines, tt.args.flags); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
