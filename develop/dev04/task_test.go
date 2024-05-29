package dev04

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	type args struct {
		dictionary []string
	}
	tests := []struct {
		name string
		args args
		want map[string][]string
	}{
		{"Default",
			args{[]string{"столик", "пятак", "слиток", "пятка", "листок", "тяпка"}},
			map[string][]string{
				"столик": {"листок", "слиток", "столик"},
				"пятак":  {"пятак", "пятка", "тяпка"},
			},
		},
		{"WithUpperCase",
			args{[]string{"сТолиК", "пятак", "сЛиток", "пяткА", "лИсток", "ТяпкА"}},
			map[string][]string{
				"столик": {"листок", "слиток", "столик"},
				"пятак":  {"пятак", "пятка", "тяпка"},
			},
		},
		{"SetWithOneWord",
			args{[]string{"сТолиК", "листок", "сЛиток", "ТяпкА"}},
			map[string][]string{
				"столик": {"листок", "слиток", "столик"},
			},
		},
		{"NoResult",
			args{[]string{"сТолиК", "ТяпкА"}},
			map[string][]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindAnagrams(tt.args.dictionary); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
