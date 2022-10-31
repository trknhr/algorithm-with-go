package main

import "testing"

func Test_add(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "add non carry",
			args: args{
				"1",
				"2",
			},
			want: "3",
		},
		{
			name: "add with carry",
			args: args{
				"66",
				"66",
			},
			want: "132",
		},
		{
			name: "add with different digits",
			args: args{
				"99999",
				"99",
			},
			want: "100098",
		},
		{
			name: "add with different digits",
			args: args{
				"1368",
				"9120",
			},
			want: "10488",
		},
		{
			name: "add with zero",
			args: args{
				"0",
				"9120",
			},
			want: "9120",
		},
		{
			name: "add with zero",
			args: args{
				"0000",
				"20",
			},
			want: "20",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddWithShift(t *testing.T) {
	type args struct {
		num1         string
		num2         string
		shiftForNum2 int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "shift plus",
			args: args{
				num1:         "1368",
				num2:         "912",
				shiftForNum2: 1,
			},
			want: "10488",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddWithShift(tt.args.num1, tt.args.num2, tt.args.shiftForNum2); got != tt.want {
				t.Errorf("AddWithShift() = %v, want %v", got, tt.want)
			}
		})
	}
}
