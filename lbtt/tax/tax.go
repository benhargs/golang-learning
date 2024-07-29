package tax

import "math"

type taxBand struct {
	BandMin    float64
	BandMax    float64
	TaxDecimal float64
}

var TwoPercentBand taxBand = taxBand{
	BandMin:    145000.00,
	BandMax:    250000.00,
	TaxDecimal: 0.02,
}

var FivePercentBand taxBand = taxBand{
	BandMin:    250000.00,
	BandMax:    325000.00,
	TaxDecimal: 0.05,
}

var TenPercentBand taxBand = taxBand{
	BandMin:    325000.00,
	BandMax:    750000.00,
	TaxDecimal: 0.1,
}

var TwelvePercentBand taxBand = taxBand{
	BandMin:    750000.00,
	BandMax:    math.MaxInt64,
	TaxDecimal: 0.12,
}

var TaxBands []taxBand = []taxBand{
	TwoPercentBand, FivePercentBand, TenPercentBand, TwelvePercentBand,
}
