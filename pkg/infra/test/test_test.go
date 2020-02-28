package test

import "testing"

func TestEqual(t *testing.T) {
	msg := Equal(true, true, "")
	if msg != "" {
		t.Errorf("Equal equal wrong. expected=%#v and got=%#v", "", msg)
	}
}

func TestEqualFailed(t *testing.T) {
	msg := Equal(true, false, "Equal")
	if msg != "Equal equal wrong. expected=false and got=true" {
		t.Errorf("Equal equal wrong. msd: %#v", msg)
	}
}
