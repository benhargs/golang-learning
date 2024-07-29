package lbtt

import (
	"lbtt/tax"
	"math"
)

func calculateTotalLBTT(houseValue float64) int64 {
	var totalTaxValue float64 = 0

	for _, taxBand := range tax.TaxBands {

		if houseValue >= taxBand.BandMin {

			taxValue := calculateTaxForBand(houseValue, taxBand.BandMax, taxBand.BandMin, taxBand.TaxDecimal)

			totalTaxValue += taxValue
		}
	}

	totalTaxValue = math.Floor(totalTaxValue)

	LBTT := int64(totalTaxValue)

	return LBTT
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
