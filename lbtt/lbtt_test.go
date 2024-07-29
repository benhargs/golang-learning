package lbtt

import (
	"testing"
)

func TestCalculateTotalLBTT(t *testing.T) {

	t.Run("Calculate total LBTT for a 15,000 property", func(t *testing.T) {
		got := calculateTotalLBTT(15000)
		var want int64 = 0

		assertTestFailMessage(t, got, want)
	})

	t.Run("Calculate total LBTT for a 145,000 property", func(t *testing.T) {
		got := calculateTotalLBTT(145000)
		var want int64 = 0

		assertTestFailMessage(t, got, want)
	})

	t.Run("Calculate total LBTT for a 249,999 property", func(t *testing.T) {
		got := calculateTotalLBTT(250000 - 1)
		var want int64 = 2099

		assertTestFailMessage(t, got, want)
	})

	t.Run("Calculate total LBTT for a 250,001 property", func(t *testing.T) {
		got := calculateTotalLBTT(250001)
		var want int64 = 2100

		assertTestFailMessage(t, got, want)
	})

	t.Run("Calculate total LBTT for a 325,001 property", func(t *testing.T) {
		got := calculateTotalLBTT(325001)
		var want int64 = 5850

		assertTestFailMessage(t, got, want)
	})

	t.Run("Calculate total LBTT for a 750001 property", func(t *testing.T) {
		got := calculateTotalLBTT(750001)
		var want int64 = 48350

		assertTestFailMessage(t, got, want)
	})

	t.Run("Calculate total LBTT for a 1,000,000 property", func(t *testing.T) {
		got := calculateTotalLBTT(1000000)
		var want int64 = 78350

		assertTestFailMessage(t, got, want)
	})
}

func assertTestFailMessage(t testing.TB, got, want int64) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
