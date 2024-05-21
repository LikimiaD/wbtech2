package main

import (
	"github.com/beevik/ntp"
	"testing"
)

func TestNTPTime(t *testing.T) {
	_, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		t.Errorf("Error while getting time: %s", err)
	}
}
