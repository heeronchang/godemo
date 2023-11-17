package test

import (
	"testing"
)

func Test_recursion(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "123",
			args: args{
				n: 3,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recursion(tt.args.n); got != tt.want {
				t.Errorf("recursion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_recursion(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		recursion(10)
	}
}

func Test_tailRecursion(t *testing.T) {
	type args struct {
		n   int
		res int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "111",
			args: args{
				n: 3,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tailRecursion(tt.args.n, tt.args.res); got != tt.want {
				t.Errorf("tailRecursion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_tailRecursion(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tailRecursion(3, 0)
	}
}
