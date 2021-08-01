package _00_init_code

import (
	"reflect"
	"testing"
)

func Test_searchRange(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{target: 1, arr: []int{1}}, want: []int{0, 0}},
		{args: args{target: 6, arr: []int{5, 7, 7, 8, 8, 10}}, want: []int{-1, -1}},
		{args: args{target: 8, arr: []int{5, 7, 7, 8, 8, 10}}, want: []int{3, 4}},
		{args: args{target: 3, arr: []int{1, 2, 3, 3, 3, 3, 3, 3, 9, 10}}, want: []int{2, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.arr, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
