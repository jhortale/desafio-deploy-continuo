package main

import "testing"

func TestSqrt(t *testing.T) {
	res := sqrt()
	if res != "Code.education always Rocks!" {
		t.Error("Erro! correto é Code.education always Rocks!", res)
	}
}
