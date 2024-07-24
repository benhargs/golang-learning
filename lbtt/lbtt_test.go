package lbtt

import (
	tax "lbtt/taxValues"
	"testing"
)

func Test2PercentTaxBand(t *testing.T) {
	var minBandValue float64 = tax.TwoPercentTaxMinBase
	var maxBandValue float64 = tax.FivePercentTaxMinBase

	t.Run("Maximum amount of tax for the 2% band", func(t *testing.T) {
		got := calculateTaxBandof2Percent(maxBandValue)
		var want float64 = 2100

		assertTestFailMessage(t, got, want)
	})

	t.Run("Minimum amount of tax for the 2% band", func(t *testing.T) {
		got := calculateTaxBandof2Percent(minBandValue + 1)
		want := 0.02

		assertTestFailMessage(t, got, want)
	})

	t.Run("No tax amount for the 2% band", func(t *testing.T) {
		got := calculateTaxBandof2Percent(minBandValue)
		want := 0.00

		assertTestFailMessage(t, got, want)
	})

	t.Run("Calculate total LBTT for a 249,999 property", func(t *testing.T) {
		got := calculateTaxBandof2Percent(maxBandValue - 1)
		want := 2099.98

		assertTestFailMessage(t, got, want)
	})

}

func Test5PercentTaxBand(t *testing.T) {
	var minBandValue float64 = tax.FivePercentTaxMinBase
	var maxBandValue float64 = tax.TenPercentTaxMinBase

	t.Run("Maximum amount of tax for the 5% band", func(t *testing.T) {
		got := calculateTaxBandof5Percent(maxBandValue)
		want := 3750.00

		assertTestFailMessage(t, got, want)
	})

	t.Run("Minimum amount of tax for 5% the band", func(t *testing.T) {
		got := calculateTaxBandof5Percent(minBandValue + 1)
		want := 0.05

		assertTestFailMessage(t, got, want)
	})

	t.Run("No tax amount for the 5% band", func(t *testing.T) {
		got := calculateTaxBandof5Percent(minBandValue)
		want := 0.00

		assertTestFailMessage(t, got, want)
	})
}

func Test10PercentTaxBand(t *testing.T) {
	var minBandValue float64 = tax.TenPercentTaxMinBase
	var maxBandValue float64 = tax.TwelvePercentTaxMinBase

	t.Run("Maximum amount of tax for the 10% band", func(t *testing.T) {
		got := calculateTaxBandOf10Percent(maxBandValue)
		want := float64(42500)

		assertTestFailMessage(t, got, want)
	})

	t.Run("Minimum amount of tax for 10% the band", func(t *testing.T) {
		got := calculateTaxBandOf10Percent(minBandValue + 1)
		want := 0.1

		assertTestFailMessage(t, got, want)
	})

	t.Run("No tax amount for the 10% band", func(t *testing.T) {
		got := calculateTaxBandOf10Percent(minBandValue)
		want := float64(0)

		assertTestFailMessage(t, got, want)
	})
}

func Test12PercentTaxBand(t *testing.T) {
	var minBandValue float64 = tax.TwelvePercentTaxMinBase
	var taxValue float64 = tax.TwelvePercentTaxValue

	t.Run("Minimum amount of tax for 10% the band", func(t *testing.T) {
		got := calculateTaxBandOf12Percent(minBandValue + 1)
		want := taxValue

		assertTestFailMessage(t, got, want)
	})

	t.Run("No tax amount for the 12% band", func(t *testing.T) {
		got := calculateTaxBandOf12Percent(minBandValue)
		want := float64(0)

		assertTestFailMessage(t, got, want)
	})
}

func TestCalculateTotalLBTT(t *testing.T) {
	var min2PercentValue float64 = tax.TwoPercentTaxMinBase
	var min5PercentValue float64 = tax.TwoPercentTaxMinBase
	var min10PercentValue float64 = tax.TwoPercentTaxMinBase
	var min12PercentValue float64 = tax.TwoPercentTaxMinBase

	t.Run("Calculate total LBTT for a 145,000 property", func(t *testing.T) {
		got := calculateTotalLBTT(min2PercentValue)
		want := float64(0)

		assertTestFailMessage(t, got, want)
	})

	t.Run("Calculate total LBTT for a 249,999 property", func(t *testing.T) {
		got := calculateTotalLBTT(min5PercentValue - 1)
		want := 2099.98

		assertTestFailMessage(t, got, want)
	})

	t.Run("Calculate total LBTT for a 250,001 property", func(t *testing.T) {
		got := calculateTotalLBTT(min5PercentValue + 1)
		want := 2100.05

		assertTestFailMessage(t, got, want)
	})

	t.Run("Calculate total LBTT for a 325,001 property", func(t *testing.T) {
		got := calculateTotalLBTT(min10PercentValue + 1)
		want := 5850.1

		assertTestFailMessage(t, got, want)
	})

	t.Run("Calculate total LBTT for a 750001 property", func(t *testing.T) {
		got := calculateTotalLBTT(min12PercentValue + 1)
		want := 48350.12

		assertTestFailMessage(t, got, want)
	})

	t.Run("Calculate total LBTT for a 1,000,000 property", func(t *testing.T) {
		got := calculateTotalLBTT(1000000)
		var want float64 = 78350

		assertTestFailMessage(t, got, want)
	})
}

func assertTestFailMessage(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
