package main

import (
	"net"
	"testing"
)

func TestIPScan(t *testing.T) {
	type args struct {
		ip net.IP
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IPScan(tt.args.ip)
		})
	}
}
