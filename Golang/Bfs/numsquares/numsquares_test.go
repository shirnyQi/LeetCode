package numsquares

import "testing"

var (
	num1  = 12 // 12 = 4 + 4 + 4
	want1 = 3

	num2  = 13 // 13 = 4 + 9
	want2 = 2
)

func Test_numSquares(t *testing.T) {
	var got int
	got = numSquares(num1)
	if want1 != got {
		t.Errorf("want:%v, got:%v", want1, got)
	}

	got = numSquares(num2)
	if want2 != got {
		t.Errorf("want:%v, got:%v", want2, got)
	}
}
