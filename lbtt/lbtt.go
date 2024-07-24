package lbtt

import (
	"fmt"
	tax "lbtt/taxValues"
)

/*
Land and Buildings Transaction Tax (LBTT) is applied to residential properties.

A return must be sent for all resendential property worth over £40,000.

The tax bands are progressive so you only pay that percentage tax for any amount into that band.

If property is worth more than £40,000 a return must be submitted, even if the tax amount is £0.

Rates:

Up to £145,000 = 0%

£145,001 to £250,000 = 2%

££250,001 to £325,000 = 5%

£325,001 to £750,000 = 10%

Over £750,000 = 12%

Example:
Value = 235,000

£90,000 in band 2, B2: 2%
So the LBTT amount is: £1,800
*/

func calculateTotalLBTT(houseValue float64) float64 {
	var (
		totalTaxValue float64 = 0

		//Calling global variables from tax package
		twoPercentTaxMin    = tax.TwoPercentTaxMinBase
		fivePercentTaxMin   = tax.FivePercentTaxMinBase
		tenPercentTaxMin    = tax.TenPercentTaxMinBase
		twelvePercentTaxMin = tax.TwelvePercentTaxMinBase
	)

	if houseValue <= twoPercentTaxMin {
		return totalTaxValue
	}

	if houseValue >= twoPercentTaxMin {
		tax2PercentBandValue := calculateTaxBandof2Percent(houseValue)
		fmt.Println(tax2PercentBandValue)
		totalTaxValue += tax2PercentBandValue
		fmt.Println(totalTaxValue)
	}

	if houseValue >= fivePercentTaxMin {
		tax5PercentBandValue := calculateTaxBandof5Percent(houseValue)

		totalTaxValue += tax5PercentBandValue
	}

	if houseValue >= tenPercentTaxMin {
		tax10PercentBandValue := calculateTaxBandOf10Percent(houseValue)

		totalTaxValue += tax10PercentBandValue
	}

	if houseValue >= twelvePercentTaxMin {
		tax12PercentBandValue := calculateTaxBandOf12Percent(houseValue)

		totalTaxValue += tax12PercentBandValue
	}

	return totalTaxValue
}

func calculateTaxBandof2Percent(houseValue float64) float64 {
	var bandMin float64 = tax.TwoPercentTaxMinBase
	var bandMax float64 = tax.FivePercentTaxMinBase
	var taxDecimal float64 = tax.TwoPercentTaxValue

	taxValue := calculateTaxForNonMaxBands(houseValue, bandMax, bandMin, taxDecimal)

	return taxValue
}

func calculateTaxBandof5Percent(houseValue float64) float64 {
	var bandMin float64 = tax.FivePercentTaxMinBase
	var bandMax float64 = tax.TenPercentTaxMinBase
	var taxDecimal float64 = tax.FivePercentTaxValue

	//Beyond Max bound
	if houseValue >= bandMax {
		taxValue := calculateMaxTaxForABound(bandMax, bandMin, taxDecimal)
		return taxValue
	}

	//Tax that is within the bounds
	if bandMin < houseValue {

		taxValue := calculateTaxWithinABound(houseValue, bandMin, taxDecimal)

		return taxValue
	}

	return 0.00
}

func calculateTaxBandOf10Percent(houseValue float64) float64 {
	var bandMin float64 = tax.TenPercentTaxMinBase
	var bandMax float64 = tax.TwelvePercentTaxMinBase
	var taxDecimal float64 = tax.TenPercentTaxValue

	if houseValue >= bandMax {
		taxValue := calculateMaxTaxForABound(bandMax, bandMin, taxDecimal)

		return taxValue
	}

	if houseValue >= bandMin {

		taxValue := calculateTaxWithinABound(houseValue, bandMin, taxDecimal)

		return taxValue
	}

	return 0.00
}

func calculateTaxBandOf12Percent(houseValue float64) float64 {
	var bandMin float64 = tax.TwelvePercentTaxMinBase
	var taxDecimal float64 = tax.TwelvePercentTaxValue

	if houseValue >= bandMin {
		taxValue := calculateTaxWithinABound(houseValue, bandMin, taxDecimal)

		return taxValue
	}

	return 0.00
}

func calculateTaxWithinABound(houseValue, bandMin, taxDecimal float64) float64 {
	bandValue := houseValue - bandMin
	taxValue := bandValue * taxDecimal

	return taxValue
}

func calculateMaxTaxForABound(bandMax, bandMin, taxDecimal float64) float64 {
	bandRange := bandMax - bandMin
	maxTaxValue := bandRange * taxDecimal

	return maxTaxValue
}

func calculateTaxForNonMaxBands(houseValue, bandMax, bandMin, taxDecimal float64) float64 {
	var taxValue float64 = 0
	var bandValue float64 = houseValue - bandMin

	if houseValue >= bandMax {

		taxValue = bandMax * taxDecimal

		return taxValue
	}

	if houseValue >= bandMin {

		taxValue = bandValue * taxDecimal

		return taxValue
	}

	return float64(0)
}
