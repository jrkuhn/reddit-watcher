package matcher

import (
	"testing"
)

var totalTests = []struct {
	in  string
	out string
}{
	{
		"[US-TX] [H] PrimeCap / CM PBT L Cherry MX Blues [W] Paypal, Local Cash",
		"PrimeCap / CM PBT L Cherry MX Blues",
	},
	{
		"[CA-ON] [H] BKE Redux Heavy, FC660C 45g Topre Domes, Leopold Keycaps Doubleshot PBT Dolch [W] PayPal ",
		"Not in the US: CA, ON",
	},
	{
		"[US-PA][H] PayPal, Local Cash [W] 65g r7+ zealios, zeal stabs r2 or newer",
		"Not for sale: 65g r7+ zealios, zeal stabs r2 or newer",
	},
	{
		"[US-FL] [H] PayPal [W] ~60 alps orange or salmon",
		"Not for sale: ~60 alps orange or salmon",
	},
	{
		"[Vendor] GMK Olivia GB MOQ Hit! - Last day! | NovelKeys Restocks | BOX Royals Update",
		"Not parsable: [Vendor] GMK Olivia GB MOQ Hit! - Last day! | NovelKeys Restocks | BOX Royals Update",
	},
	{
		"[MY][H] Fuguthulus, matrix abels, artisans [W] Trades, Paypal ",
		"Not parsable: [MY][H] Fuguthulus, matrix abels, artisans [W] Trades, Paypal ",
	},
}

func TestAll(t *testing.T) {
	for _, tt := range totalTests {
		out, err := FindValidSale(tt.in)

		if err != nil {
			out = err.Error()
		}
		if out != tt.out {
			t.Errorf("Expected %q, got %q", tt.out, out)
		}
	}
}