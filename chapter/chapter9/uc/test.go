package uc

import "testing"

type ucTest struct {
	in, out string
}

var unTests = []ucTest{
	{"abc", "ABC"},
	{"cvo-az", "CVO-AZ"},
	{"antwerp", "ANTWERP"},
}

func TestUC(t *testing.T) {
	for _, ut := range unTests {
		uc := UpperCase(ut.in)
		if uc != ut.out {
			t.Errorf("UpperCase(%s) = %s, must be %s", ut.in, uc, ut.out)
		}
	}
}
