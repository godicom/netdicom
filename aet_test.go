package netdicom

import (
	"testing"
)

func TestSetAEPositive(t *testing.T) {
	cases := []struct {
		aet, addr string
		timeout   uint
	}{
		{"aet1", "localhost:11112", 10},
		{"aet2", "127.0.0.1:11112", 100},
	}

	for _, c := range cases {
		ae, err := NewAE(AET(c.aet), Timeout(c.timeout), Address(c.addr))
		if err != nil {
			t.Fatalf("Failed to create AE instance with error: %q", err)
		}
		aet := ae.AET()
		if aet != c.aet {
			t.Errorf("Failed to set AET %q, got %q", c.aet, aet)
		}
		addr := ae.Address()
		if addr != c.addr {
			t.Errorf("Failed to set Address %q, got %q", c.addr, addr)
		}
		timeout := ae.Timeout()
		if timeout != c.timeout {
			t.Errorf("Failed to set Timeout %q, got %q", c.timeout, timeout)
		}
	}
}

func TestAEDefault(t *testing.T) {
	ae, err := NewAE()
	if err != nil {
		t.Fatalf("Failed to create AE instance with error: %q", err)
	}

	aet := ae.AET()
	if aet != "aet" {
		t.Errorf("Incorrect default AET: %q", aet)
	}

	addr := ae.Address()
	if addr != "" {
		t.Errorf("Default local address should be empty, got: %q", addr)
	}

	timeout := ae.Timeout()
	if timeout != 0 {
		t.Errorf("Current default timeout expected to be 0, got: %q", timeout)
	}
}
