package slicestring

import "testing"

func TestContains(t *testing.T) {
	ss := []string{"abc", "cde", "fgh"}
	k := "fgh"

	obtained := Contains(ss, k)
	if !obtained {
		t.Errorf("Contains() should return true when []string contains keyword: obtained=%+v, []string=%+v, keyword=%+v", obtained, ss, k)
	}

}
