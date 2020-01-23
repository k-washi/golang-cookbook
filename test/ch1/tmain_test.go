package ch1

import (
	"testing"
)

func TestSimpleTestFunc(t *testing.T) {
	/*
			go test -v  ./test/ch1/...
		=== RUN   TestSimpleTestFunc
		--- PASS: TestSimpleTestFunc (0.00s)
		    tmain_test.go:12: SimpleTestFunc() => nil
		    tmain_test.go:28: SimpleTestFunc(1, 2, 3) -> []int{1, 2, 3}
		    tmain_test.go:43: SimpleTestFunc([]int{1, 2, 3}...) -> []int{1, 2, 3}
		PASS
		ok      github.com/k-washi/golang-cookbook/test/ch1     0.010s
	*/
	//引数なし
	if v := SimpleTestFunc(); v != nil {
		t.Error("引数なしのためnilを返すべき", nil)
	} else {
		t.Log("SimpleTestFunc() => nil")
	}

	//引数あり
	v := SimpleTestFunc(1, 2, 3)
	expected := []int{1, 2, 3}
	isErr := false
	for i := 0; i < 3; i++ {
		if v[i] != expected[i] {
			isErr = true
			break
		}
	}
	if isErr {
		t.Error("args v != ]int{1, 2, 3}", v)
	} else {
		t.Log("SimpleTestFunc(1, 2, 3) -> []int{1, 2, 3}")
	}

	//sliceのテスト
	v = SimpleTestFunc(expected...)
	isErr = false
	for i := 0; i < 3; i++ {
		if v[i] != expected[i] {
			isErr = true
			break
		}
	}
	if isErr {
		t.Error("slice v != ]int{1, 2, 3}", v)
	} else {
		t.Log("SimpleTestFunc([]int{1, 2, 3}...) -> []int{1, 2, 3}")
	}

}

func TestAddInt(t *testing.T) {
	tCase := []struct {
		Name string
		Val  []int
		Exp  int
	}{
		{"addInt() -> 0", []int{}, 0},
		{"addInt([]int{5, 10, 20}) -> 35", []int{5, 10, 20}, 35},
	}

	for _, tc := range tCase {
		/*
					=== RUN   TestAddInt
			=== RUN   TestAddInt/addInt()_->_0
			=== RUN   TestAddInt/addInt([]int{5,_10,_20})_->_35
			--- PASS: TestAddInt (0.00s)
			    --- PASS: TestAddInt/addInt()_->_0 (0.00s)
			        tmain_test.go:74: 0==0
			    --- PASS: TestAddInt/addInt([]int{5,_10,_20})_->_35 (0.00s)
							tmain_test.go:74: 35==35
		*/
		t.Run(tc.Name, func(t *testing.T) {
			sum := AddInt(tc.Val...)
			if sum != tc.Exp {
				t.Errorf("%d != %d", sum, tc.Exp)
			} else {
				t.Logf("%d==%d", sum, tc.Exp)
			}
		})
	}

}
