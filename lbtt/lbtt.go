package main

import "lbtt/taxValues/taxValues2024/tax"

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

func nilTaxBelow145K(houseValue float64) float64 {
	return float64(0)
}

func calculateTaxBandof2Percent(houseValue float64) float64 {
	var bandMin float64 = tax.TwoPercentTaxBase
	var bandMax float64 = tax.FivePercentTaxBase
	var taxDecimal float64 = tax.TwoPercentTaxValue

	if houseValue >= bandMax {

		taxValue := calculateMaxTaxForABound(bandMax, bandMin, taxDecimal)

		return taxValue
	}

	if bandMin < houseValue {
		taxValue := calculateTaxWithinABound(houseValue, taxDecimal, bandMin)

		return float64(taxValue)
	}

	return float64(0)
}

func calculateTaxBandof5Percent(houseValue float64) float64 {
	var bandMin float64 = tax.FivePercentTaxBase
	var bandMax float64 = tax.TenPercentTaxMinBase
	var taxDecimal float64 = tax.FivePercentTaxValue

	//Beyond Max bound
	if houseValue >= bandMax {
		taxValue := calculateMaxTaxForABound(bandMax, bandMin, taxDecimal)
		return taxValue
	}

	//Tax that is within the bounds
	if bandMin < houseValue {

		taxValue := calculateMaxTaxForABound(bandMax, bandMin, taxDecimal)

		return taxValue
	}

	return float64(0)
}

func calculateTaxBandOf10Percent(houseValue float64) float64 {
	var bandMin float64 = tax.TenPercentTaxBase
	var bandMax float64 = tax.TwelvePercentTaxBase
	var taxDecimal float64 = tax.TenPercentTaxValue

	if houseValue >= bandMax {
		taxValue := calculateMaxTaxForABound(bandMax, bandMin, taxDecimal)

		return taxValue
	}

	if houseValue > bandMin {

		taxValue := calculateMaxTaxForABound(bandMax, bandMin, taxDecimal)

		return taxValue
	}

	return float64(0)
}

func calculateTaxBandOf12Percent(houseValue float64) float64 {
	bandBase := float64(750000)
	taxBase := 0.12

	if houseValue > bandBase {
		taxValue := calculateTaxWithinABound(houseValue, taxBase, bandBase)

		return taxValue
	}

	return float64(0)
}

func calculateTaxWithinABound(houseValue, taxBase, bandBase float64) float64 {
	bandValue := houseValue - bandBase
	taxValue := bandValue * taxBase

	return taxValue
}

func calculateMaxTaxForABound(bandMax, bandMin, taxDecimal float64) float64 {
	bandRange := bandMax - bandMin
	maxTaxValue := bandRange * taxDecimal

	return maxTaxValue
}

func calculateTotalLBTT(houseValue float64) float64 {
	var totalTaxValue float64 = 0

	//Calling global variables from tax package
	var twoPercentTaxMin = tax.TwoPercentTaxMinBase
	var fivePercentTaxMin = tax.FivePercentTaxMinBase
	var tenPercentTaxMin = tax.TenPercentTaxMinBase
	var twelvePercentTaxMin = tax.TwelvePercentTaxMinBase

	if houseValue < twoPercentTaxMin {
		totalTaxValue = nilTaxBelow145K(houseValue)

		return totalTaxValue
	}

	if houseValue > twoPercentTaxMin {
		tax2PercentBandValue := calculateTaxBandof2Percent(houseValue)

		totalTaxValue += tax2PercentBandValue
	}

	if houseValue > fivePercentTaxMin {
		tax5PercentBandValue := calculateTaxBandof5Percent(houseValue)

		totalTaxValue += tax5PercentBandValue
	}

	if houseValue > tenPercentTaxMin {
		tax10PercentBandValue := calculateTaxBandOf10Percent(houseValue)

		totalTaxValue += tax10PercentBandValue
	}

	if houseValue > twelvePercentTaxMin {
		tax12PercentBandValue := calculateTaxBandOf12Percent(houseValue)

		totalTaxValue += tax12PercentBandValue
	}

	return totalTaxValue
}
