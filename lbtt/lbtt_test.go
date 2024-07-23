package main

import "testing"

func TestTaxFreeBand(t *testing.T) {
	got := nilTaxBelow145K(144000.00)
	want := float64(0)

	assertTestFailMessage(t, got, want)
}

func Test2PercentTaxBand(t *testing.T) {

	t.Run("Maximum amount of tax for the 2% band", func(t *testing.T) {
		got := calculateTaxBandof2Percent(1000000)
		var want float64 = 2100

		assertTestFailMessage(t, got, want)
	})

	t.Run("Minimum amount of tax for the 2% band", func(t *testing.T) {
		got := calculateTaxBandof2Percent(145001)
		want := 0.02

		assertTestFailMessage(t, got, want)
	})

	t.Run("No tax amount for the 2% band", func(t *testing.T) {
		got := calculateTaxBandof2Percent(145000)
		want := float64(0)

		assertTestFailMessage(t, got, want)
	})

}

func Test5PercentTaxBand(t *testing.T) {

	t.Run("Maximum amount of tax for the 5% band", func(t *testing.T) {
		got := calculateTaxBandof5Percent(1000000)
		want := 3750.00

		assertTestFailMessage(t, got, want)
	})

	t.Run("Minimum amount of tax for 5% the band", func(t *testing.T) {
		got := calculateTaxBandof5Percent(250001)
		want := 0.05

		assertTestFailMessage(t, got, want)
	})

	t.Run("No tax amount for the 5% band", func(t *testing.T) {
		got := calculateTaxBandof5Percent(250000)
		want := float64(0)

		assertTestFailMessage(t, got, want)
	})
}

func Test10PercentTaxBand(t *testing.T) {

	t.Run("Maximum amount of tax for the 10% band", func(t *testing.T) {
		got := calculateTaxBandOf10Percent(1000000)
		want := float64(42500)

		assertTestFailMessage(t, got, want)
	})

	t.Run("Minimum amount of tax for 10% the band", func(t *testing.T) {
		got := calculateTaxBandOf10Percent(325001)
		want := 0.1

		assertTestFailMessage(t, got, want)
	})

	t.Run("No tax amount for the 10% band", func(t *testing.T) {
		got := calculateTaxBandOf10Percent(250000)
		want := float64(0)

		assertTestFailMessage(t, got, want)
	})
}

func Test12PercentTaxBand(t *testing.T) {

	t.Run("Minimum amount of tax for 10% the band", func(t *testing.T) {
		got := calculateTaxBandOf12Percent(750001)
		want := 0.12

		assertTestFailMessage(t, got, want)
	})

	t.Run("No tax amount for the 12% band", func(t *testing.T) {
		got := calculateTaxBandOf12Percent(750000)
		want := float64(0)

		assertTestFailMessage(t, got, want)
	})
}

func TestCalculateTotalLBTT(t *testing.T) {
	t.Run("Calculate total LBTT for a 145000 property", func(t *testing.T) {
		got := calculateTotalLBTT(145000)
		want := float64(0)

		assertTestFailMessage(t, got, want)
	})

	t.Run("Calculate total LBTT for a 249999 property", func(t *testing.T) {
		got := calculateTotalLBTT(249999)
		want := 2099.98

		assertTestFailMessage(t, got, want)
	})

	t.Run("Calculate total LBTT for a 250001 property", func(t *testing.T) {
		got := calculateTotalLBTT(250001)
		want := 2100.05

		assertTestFailMessage(t, got, want)
	})

	t.Run("Calculate total LBTT for a 250001 property", func(t *testing.T) {
		got := calculateTotalLBTT(325001)
		want := 5850.1

		assertTestFailMessage(t, got, want)
	})

	t.Run("Calculate total LBTT for a 750001 property", func(t *testing.T) {
		got := calculateTotalLBTT(750001)
		want := 48350.12

		assertTestFailMessage(t, got, want)
	})

	t.Run("Calculate total LBTT for a 1,000,000 property", func(t *testing.T) {
		got := calculateTotalLBTT(1000000)
		want := float64(78350)

		assertTestFailMessage(t, got, want)
	})
}

func assertTestFailMessage(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
