package lbtt

import (
	tax "lbtt/taxValues"
	"math"
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

£90,000 into band 2, B2: 2%
So the LBTT amount is: £1,800
*/

func calculateTotalLBTT(houseValue float64) float64 {
	var totalTaxValue float64 = 0

	//Calling variables from tax package
	var twoPercentTaxMin float64 = tax.TwoPercentTaxMinBase
	var fivePercentTaxMin float64 = tax.FivePercentTaxMinBase
	var tenPercentTaxMin float64 = tax.TenPercentTaxMinBase
	var twelvePercentTaxMin float64 = tax.TwelvePercentTaxMinBase

	if houseValue >= twoPercentTaxMin {

		tax2PercentBandValue := calculateTaxBandof2Percent(houseValue)

		totalTaxValue += tax2PercentBandValue
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

	taxValue := calculateTaxForNonMaxBand(houseValue, bandMax, bandMin, taxDecimal)

	return taxValue
}

func calculateTaxBandof5Percent(houseValue float64) float64 {
	var bandMin float64 = tax.FivePercentTaxMinBase
	var bandMax float64 = tax.TenPercentTaxMinBase
	var taxDecimal float64 = tax.FivePercentTaxValue

	taxValue := calculateTaxForNonMaxBand(houseValue, bandMax, bandMin, taxDecimal)

	return taxValue
}

func calculateTaxBandOf10Percent(houseValue float64) float64 {
	var bandMin float64 = tax.TenPercentTaxMinBase
	var bandMax float64 = tax.TwelvePercentTaxMinBase
	var taxDecimal float64 = tax.TenPercentTaxValue

	taxValue := calculateTaxForNonMaxBand(houseValue, bandMax, bandMin, taxDecimal)

	return taxValue
}

func calculateTaxBandOf12Percent(houseValue float64) float64 {
	var bandMin float64 = tax.TwelvePercentTaxMinBase
	var taxDecimal float64 = tax.TwelvePercentTaxValue

	bandValue := (houseValue - bandMin)
	taxValue := (bandValue * taxDecimal)

	return taxValue
}

func calculateTaxForNonMaxBand(houseValue, bandMax, bandMin, taxDecimal float64) float64 {
	var bandValue float64 = houseValue - bandMin
	var bandRange float64 = bandMax - bandMin

	if houseValue > bandMax {

		taxValue := (bandRange * taxDecimal)

		return taxValue
	}

	if houseValue >= bandMin {

		taxValue := (bandValue * taxDecimal)

		taxValue = math.Round(taxValue*100) / 100

		return taxValue
	}

	return float64(0)
}
