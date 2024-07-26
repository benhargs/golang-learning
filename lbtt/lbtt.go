package lbtt

import (
	"lbtt/tax"
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

	for _, taxBand := range tax.TaxBands {

		if houseValue >= taxBand.BandMin {

			taxValue := calculateTaxForBand(houseValue, taxBand.BandMax, taxBand.BandMin, taxBand.TaxDecimal)

			totalTaxValue += taxValue
		}
	}

	totalTaxValue = math.Floor(totalTaxValue)

	return totalTaxValue
}

func calculateTaxForBand(houseValue, bandMax, bandMin, taxDecimal float64) float64 {
	var bandValue float64 = houseValue - bandMin
	var bandRange float64 = bandMax - bandMin

	if houseValue > bandMax {

		taxValue := (bandRange * taxDecimal)

		return taxValue
	}

	if houseValue >= bandMin {

		taxValue := (bandValue * taxDecimal)

		return taxValue
	}

	return float64(0)
}
