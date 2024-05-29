package dev02

import (
	"testing"
)

func TestUnpackString(t *testing.T) {
	type args struct {
		source string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"WithoutRepeat", args{source: "abcd"}, "abcd", false},
		{"Repeated", args{source: "a4bc2d5e"}, "aaaabccddddde", false},
		{"RepeatMoreThanNine", args{source: "b10"}, "bbbbbbbbbb", false},
		{"OneCharacter", args{source: "z"}, "z", false},
		{"OneCharacterRepeated", args{source: "z2"}, "zz", false},
		{"EmptyString", args{source: ""}, "", false},

		{"IncorrectString", args{source: "45"}, "", true},

		{"WithEscape", args{source: `qwe\4\5`}, "qwe45", false},
		{"WithEscapeRepeated", args{source: `qwe\45`}, "qwe44444", false},
		{"WithEscapedEscapeRepeated", args{source: `qwe\\5`}, `qwe\\\\\`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnpackString(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnpackString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UnpackString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
