package main

import "testing"

func TestPrintCurrentTime(t *testing.T) {
	type args struct {
		server string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"DefaultTest", args{server: "0.beevik-ntp.pool.ntp.org"}, false},
		{"UnknownServerTest", args{server: "google.com"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PrintCurrentTime(tt.args.server); (err != nil) != tt.wantErr {
				t.Errorf("PrintCurrentTime() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
