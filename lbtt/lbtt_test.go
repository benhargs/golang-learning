package lbtt

import (
	tax "lbtt/taxValues"
	"testing"
)

func TestTaxBandswithMaxBoundaryCases(t *testing.T) {

	//2 percent tax values:
	var twoPercentTaxMinBase float64 = tax.TwoPercentTaxMinBase
	var twoPercentTaxValue float64 = tax.TwoPercentTaxValue

	// 5 percent tax values:
	var fivePercentTaxMinBase float64 = tax.FivePercentTaxMinBase
	var fivePercentTaxValue float64 = tax.FivePercentTaxValue

	// 10 percent tax values:
	var tenPercentTaxMinBase float64 = tax.TenPercentTaxMinBase
	var tenPercentTaxValue float64 = tax.TenPercentTaxValue
	var tenPercentTaxMax float64 = tax.TwelvePercentTaxMinBase

	taxBand := []struct {
		name       string
		percentage int
		houseValue float64
		want       float64
	}{
		//2 percent tests
		{"No Tax", 2, twoPercentTaxMinBase, 0},
		{"Minimum amount of Tax in bound", 2, twoPercentTaxMinBase + 1, twoPercentTaxValue},
		{"Maximum amount of property value -1 in bound", 2, fivePercentTaxMinBase - 1, 2099.98},
		{"Maximum amount of within max band of Tax", 2, fivePercentTaxMinBase, 2100},
		{"No more Tax beyond max bound", 2, fivePercentTaxMinBase + 1, 2100},

		// five percent tests
		{"No Tax", 5, fivePercentTaxMinBase, 0},
		{"Minimum amount of Tax", 5, fivePercentTaxMinBase + 1, fivePercentTaxValue},
		{"Maximum amount property value -1 in bound", 5, tenPercentTaxMinBase - 1, 3749.95},
		{"Maximum amount of within band Tax", 5, tenPercentTaxMinBase, 3750},
		{"Maximum amount of beyond band Tax", 5, tenPercentTaxMinBase + 1, 3750},

		//10 percent tests
		{"No Tax", 10, fivePercentTaxMinBase, 0},
		{"Minimum amount of Tax", 10, tenPercentTaxMinBase + 1, tenPercentTaxValue},
		{"Maximum amount property value -1 in bound", 10, tenPercentTaxMax - 1, 42499.9},
		{"Maximum amount of within band Tax", 10, tenPercentTaxMax, 42500},
		{"Maximum amount of beyond band Tax", 10, tenPercentTaxMax + 1, 42500},
	}

	for _, testCase := range taxBand {
		want := testCase.want

		if testCase.percentage == 2 {

			t.Run(testCase.name, func(t *testing.T) {
				got := calculateTaxBandof2Percent(testCase.houseValue)

				assertTestFailMessage(t, got, want)
			})

		}

		if testCase.percentage == 5 {

			t.Run(testCase.name, func(t *testing.T) {
				got := calculateTaxBandof5Percent(testCase.houseValue)

				assertTestFailMessage(t, got, want)
			})

		}

		if testCase.percentage == 10 {

			t.Run(testCase.name, func(t *testing.T) {
				got := calculateTaxBandOf10Percent(testCase.houseValue)

				assertTestFailMessage(t, got, want)
			})

		}

	}
}
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

	t.Run("Calculate LBTT for a Max -1, of a 2% property", func(t *testing.T) {
		got := calculateTaxBandof2Percent(maxBandValue - 1)
		want := 2099.98

		assertTestFailMessage(t, got, want)
	})

}

func Test5PercentTaxBand(t *testing.T) {
	var minBandValue float64 = tax.FivePercentTaxMinBase
	var maxBandValue float64 = tax.TenPercentTaxMinBase
	var taxDecimal float64 = tax.FivePercentTaxValue

	t.Run("Maximum amount of tax for the 5% band", func(t *testing.T) {
		got := calculateTaxBandof5Percent(maxBandValue)
		want := 3750.00

		assertTestFailMessage(t, got, want)
	})

	t.Run("Minimum amount of tax for 5% the band", func(t *testing.T) {
		got := calculateTaxBandof5Percent(minBandValue + 1)
		want := taxDecimal

		assertTestFailMessage(t, got, want)
	})

	t.Run("No tax amount for the 5% band", func(t *testing.T) {
		got := calculateTaxBandof5Percent(minBandValue)
		want := 0.00

		assertTestFailMessage(t, got, want)
	})

	t.Run("Max value -1 amount of tax for the 5% band", func(t *testing.T) {
		got := calculateTaxBandof5Percent(maxBandValue - 1)
		want := float64(3749.95)

		assertTestFailMessage(t, got, want)
	})
}

func Test10PercentTaxBand(t *testing.T) {
	var minBandValue float64 = tax.TenPercentTaxMinBase
	var maxBandValue float64 = tax.TwelvePercentTaxMinBase
	var taxDecimal float64 = tax.TenPercentTaxValue

	t.Run("Maximum amount of tax for the 10% band", func(t *testing.T) {
		got := calculateTaxBandOf10Percent(maxBandValue)
		want := float64(42500)

		assertTestFailMessage(t, got, want)
	})

	t.Run("Max value -1 amount of tax for the 10% band", func(t *testing.T) {
		got := calculateTaxBandOf10Percent(maxBandValue - 1)
		want := float64(42499.9)

		assertTestFailMessage(t, got, want)
	})

	t.Run("Minimum amount of tax for the 10% band", func(t *testing.T) {
		got := calculateTaxBandOf10Percent(minBandValue + 1)
		want := taxDecimal

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
	var taxDecimal float64 = tax.TwelvePercentTaxValue

	t.Run("Minimum amount of tax for 12% the band", func(t *testing.T) {
		got := calculateTaxBandOf12Percent(minBandValue + 1)
		want := taxDecimal

		assertTestFailMessage(t, got, want)
	})

	t.Run("No tax amount for the 12% band", func(t *testing.T) {
		got := calculateTaxBandOf12Percent(minBandValue)
		want := float64(0)

		assertTestFailMessage(t, got, want)
	})

	t.Run("£1,000,000 property, for the 12% band", func(t *testing.T) {
		got := calculateTaxBandOf12Percent(minBandValue + 250000)
		want := float64(30000)

		assertTestFailMessage(t, got, want)
	})
}

func TestCalculateTotalLBTT(t *testing.T) {
	var min2PercentValue float64 = tax.TwoPercentTaxMinBase
	var min5PercentValue float64 = tax.FivePercentTaxMinBase
	var min10PercentValue float64 = tax.TenPercentTaxMinBase
	var min12PercentValue float64 = tax.TwelvePercentTaxMinBase

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
