/*
   @Time : 2020-03-19 14:19
   @Author : liuoneice
   @File : demo1
   @Software: GoLand
*/
package bishiti

import (
	"reflect"
	"testing"
)

func Test_fib1(t *testing.T) {
	type args struct {
		i int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				i: 0,
			},
			want: 0,
		}, {
			name: "case2",
			args: args{
				i: 1,
			},
			want: 1,
		}, {
			name: "case3",
			args: args{
				i: 2,
			},
			want: 1,
		}, {
			name: "case4",
			args: args{
				i: 3,
			},
			want: 2,
		}, {
			name: "case5",
			args: args{
				i: 4,
			},
			want: 3,
		}, {
			name: "case6",
			args: args{
				i: 5,
			},
			want: 5,
		}, {
			name: "case7",
			args: args{
				i: 15,
			},
			want: 610,
		}, {
			name: "case8",
			args: args{
				i: 50,
			},
			want: 12586269025,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib1(tt.args.i); got != tt.want {
				t.Errorf("fib1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fib2(t *testing.T) {
	type args struct {
		i int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				i: 0,
			},
			want: "0",
		}, {
			name: "case2",
			args: args{
				i: 1,
			},
			want: "1",
		}, {
			name: "case3",
			args: args{
				i: 2,
			},
			want: "1",
		}, {
			name: "case4",
			args: args{
				i: 3,
			},
			want: "2",
		}, {
			name: "case5",
			args: args{
				i: 4,
			},
			want: "3",
		}, {
			name: "case6",
			args: args{
				i: 4,
			},
			want: "3",
		}, {
			name: "case7",
			args: args{
				i: 15,
			},
			want: "610",
		}, {
			name: "case8",
			args: args{
				i: 150,
			},
			want: "9969216677189303386214405760200",
		}, {
			name: "case9",
			args: args{
				i: 149,
			},
			want: "6161314747715278029583501626149",
		}, {
			name: "case10",
			args: args{
				i: 148,
			},
			want: "3807901929474025356630904134051",
		}, {
			name: "case11",
			args: args{
				i: 147,
			},
			want: "2353412818241252672952597492098",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib2(tt.args.i); !reflect.DeepEqual(got.String(), tt.want) {
				t.Errorf("fib2(%d) = %v, want %v", tt.args.i, got, tt.want)
			}
		})
	}
}
