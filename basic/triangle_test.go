package main

import "testing"

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{30000, 40000, 50000},
	}
	for _, tt := range tests {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calctriangle(%d,%d);"+
				"got %d;expected %d",
				tt.a, tt.b, actual, tt.c)

		}
	}
}

/**
性能测试
 */
func BenchmarkTriangle(b *testing.B) {
	a:=30000
	b1:=40000
	c:=50000
	for i:=0;i<b.N;i++{
		if actual := calcTriangle(a, b1); actual != c {
			b.Errorf("calctriangle(%d,%d);"+
				"got %d;expected %d",
				a, b1, actual, c)

	}
	}

}
