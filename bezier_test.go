package main

import "testing"

func TestPowers(t *testing.T) {
	testCases := []float64{0.0, 0.25, 0.5, 0.75, 1.0}
	for _, tc := range testCases {
		t.Log(Power(tc))
	}
}

func TestT(t *testing.T) {
	t.Log(Mb)
	t.Log(Mb.T())
}
