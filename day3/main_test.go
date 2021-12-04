package main

import (
	"reflect"
	"testing"
)

var testdata = []int64{
	0b00100,
	0b11110,
	0b10110,
	0b10111,
	0b10101,
	0b01111,
	0b00111,
	0b11100,
	0b10000,
	0b11001,
	0b00010,
	0b01010,
}

var testkeep1 = []int64{
	0b11110,
	0b10110,
	0b10111,
	0b10101,
	0b11100,
	0b10000,
	0b11001,
}

var testkeep2 = []int64{
	0b11110,
	0b01111,
	0b11100,
	0b11001,
	0b01010,
}

var testkeep3 = []int64{
	0b00100,
	0b11110,
	0b10110,
	0b11100,
	0b10000,
	0b00010,
	0b01010,
}

func Test_getonescount(t *testing.T) {
	type args struct {
		data []int64
		bit  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{data: testdata, bit: 5}, 7},
		{"test2", args{data: testdata, bit: 4}, 5},
		{"test3", args{data: testdata, bit: 1}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getonescount(tt.args.data, tt.args.bit); got != tt.want {
				t.Errorf("getonescount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_keep(t *testing.T) {
	type args struct {
		data  []int64
		bit   int
		value int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{"test1", args{data: testdata, bit: 5, value: 1}, testkeep1},
		{"test2", args{data: testdata, bit: 4, value: 1}, testkeep2},
		{"test3", args{data: testdata, bit: 1, value: 0}, testkeep3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := keep(tt.args.data, tt.args.bit, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("keep() = %v, want %v", got, tt.want)
			}
		})
	}
}
