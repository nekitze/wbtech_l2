package main

import "testing"

func Test_cutLine(t *testing.T) {
	type args struct {
		line      string
		delimiter string
		fields    int
		separated bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Default", args{
			line:      "data1 data2 data3 data4",
			fields:    1,
			delimiter: " ",
		},
			"data1",
		},
		{"Delimiter", args{
			line:      "data1,data2,data3,data4",
			fields:    1,
			delimiter: ",",
		},
			"data1",
		},
		{"Separated", args{
			line:      "data1data2data3data4",
			fields:    1,
			delimiter: ",",
			separated: true,
		},
			"",
		},
		{"Limit", args{
			line:      "data1,data2,data3,data4",
			fields:    5,
			delimiter: ",",
		},
			"data1,data2,data3,data4",
		},
		{"SomeField", args{
			line:      "data1,data2,data3,data4",
			fields:    3,
			delimiter: ",",
		},
			"data3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cutLine(tt.args.line, tt.args.delimiter, tt.args.fields, tt.args.separated); got != tt.want {
				t.Errorf("cutLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
